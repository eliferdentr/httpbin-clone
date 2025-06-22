package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	utils "httbinclone-eliferden.com/utils"
)

  var nonceMap = make(map[string]time.Time )

//verifies basic authentication and returns 200 if successful
//Prompts the user for authorization using HTTP Basic Auth.
//basic auth: Basic Authentication relies on the username and password being
//combined, encoded with Base64, and included on the server in the
// Authorization HTTP header.
// 1- A client makes a request to a protected server.
//2- The server sends a 401 Unauthorized response and
// a WWW-Authenticate: Basic realm="Realm Name" header to
// indicate that the resource requires authentication.
// The "realm"(yetki alanı) is typically a description of the protected realm
// and appears in the login window that browsers display to the user.
//3-The client asks the user for a username and password (or retrieves them from where it is stored)
//4-It combines the username and password in the format username:password.
//5-It encodes this combined string in Base64.
//6-It sends a new request to the same resource by appending the encoded string to an HTTP header in the format Authorization: Basic <encoded_string>.
//7-The server receives the Authorization header, decodes it in Base64, extracts the username and password, and compares it to records in its own system.
//8-If it matches, it returns the requested resource with a 200 OK.
//9-If it doesn't match, it sends back a 401 Unauthorized response.

func VerifyBasicAuth(context *gin.Context) {
	header, err := getAuthorizationHeader(context)
	response := utils.BuildResponse(context, nil)
	if err != nil {
		response["error"] = err.Error()
		context.Header("WWW-Authenticate", `Basic realm="Access to the site"`)
		context.JSON(http.StatusUnauthorized, response)
		return
	}

	// Basic auth validation
	payloadParts, err := getBasicAuthParts(header)
	if err != nil {
		response["error"] = "Failed to parse Basic auth header: " + err.Error()
		context.Header("WWW-Authenticate", `Basic realm="Access to the site"`)
		context.JSON(http.StatusUnauthorized, response)
		return
	}

	if len(payloadParts) != 2 {
		response["error"] = "Failed to parse Basic auth header"
		context.Header("WWW-Authenticate", `Basic realm="Access to the site"`)
		context.JSON(http.StatusUnauthorized, response)
		return
	}

	userParam := context.Param("user")
	passParam := context.Param("password")

	if payloadParts[0] != userParam || payloadParts[1] != passParam {
		response["error"] = "Credentials don't match URL params"
		context.Header("WWW-Authenticate", `Basic realm="Access to the site"`)
		context.JSON(http.StatusUnauthorized, response)
		return
	}

	response["message"] = "Authentication successful"
	response["authenticated"] = true
	response["user"] = userParam
	context.JSON(http.StatusOK, response)
}

func getBasicAuthParts(header string) ([]string, error) {
	authParts := strings.SplitN(header, " ", 2)
	if len(authParts) != 2 || authParts[0] != "Basic" {
		return []string{}, fmt.Errorf("Invalid Basic Auth Header Format!")
	}

	payload, err := base64.StdEncoding.DecodeString(authParts[1])
	if err != nil {
		return nil, fmt.Errorf("base64 decoding error: %w", err)
	}

	parts := strings.SplitN(string(payload), ":", 2)

	if len(parts) != 2 {
		return nil, fmt.Errorf("authorization format is invalid:")
	}
	return parts, nil
}

func VerifyBearerAuth(context *gin.Context) {
	response := utils.BuildResponse(context, nil)
	header, err := getAuthorizationHeader(context)

	if err != nil {
		response["error"] = err.Error()
		context.Header("WWW-Authenticate", `Basic realm="Access to the site"`)
		context.JSON(http.StatusUnauthorized, response)
		return
	}

	//split as "Bearer TOKEN"
	parts := strings.SplitN(header, " ", 2)

	if len(parts) != 2 || parts[0] != "Bearer" {
		response["error"] = "Authorization must be Bearer {token}"
		context.Header("WWW-Authenticate", `Bearer realm="Access to the protected area"`)
		context.JSON(http.StatusUnauthorized, response)
		return

	}

	token := parts[1]

	response["token"] = token
	response["authenticated"] = true
	response["message"] = "Bearer authentication successful"
	context.JSON(http.StatusOK, response)

}

func VerifyDigestAuthTypes(context *gin.Context) {
	// önce parametreleri al.
	// Çünkü hem Meydan Okuma (challenge) hem de Doğrulama (verification)
	// fazında bunlara ihtiyacımız var.
    // --- 1. Parametreleri Al ---
    qopParam := context.Param("qop")
    userParam := context.Param("user")
    passwdParam := context.Param("passwd")
    algorithmParam := context.Param("algorithm")
    staleAfterParam := context.Param("stale_after")

    response := utils.BuildResponse(context, nil)
    header, err := getAuthorizationHeader(context)
    realm := fmt.Sprintf("%s@eliferden.com", userParam)

    // --- 2. Meydan Okuma mı? Doğrulama mı? ---
    if err != nil {
        // --- MEYDAN OKUMA FAZI (Authorization başlığı YOK) ---
        nonce := generateSimpleNonce()

        // YENİ ADIM: Nonce'ı oluşturulma zamanıyla birlikte map'e kaydet.
        nonceMap[nonce] = time.Now()

        algoForChallenge := algorithmParam
        if algoForChallenge == "" {
            algoForChallenge = "MD5"
        }

        wwwAuthValue := fmt.Sprintf(`Digest realm="%s", qop="%s", nonce="%s", algorithm="%s"`,
            realm, qopParam, nonce, algoForChallenge)

        context.Header("WWW-Authenticate", wwwAuthValue)
        context.JSON(http.StatusUnauthorized, response)
        return
    }

    // --- DOĞRULAMA FAZI (Authorization başlığı VAR) ---
    authParams := parseDigestAuthHeader(header)
    if authParams == nil {
        // Hatalı format
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization header format"})
        return
    }

    clientNonce := authParams["nonce"]

    // --- YENİ ADIM: STALE KONTROLÜ ---
    // Bu kontrol, imza hesaplamasından ÖNCE yapılmalı.
    if staleAfterParam != "" {
        creationTime, nonceExists := nonceMap[clientNonce]
        if !nonceExists {
            // İstemci, bizim üretmediğimiz veya çoktan sildiğimiz bir nonce gönderdi.
            context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or unknown nonce"})
            return
        }

        // URL'den gelen süreyi parse et
        staleDuration, parseErr := time.ParseDuration(staleAfterParam)
        if parseErr != nil {
            context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stale_after format", "details": parseErr.Error()})
            return
        }

        // Nonce'ın yaşını hesapla
        nonceAge := time.Since(creationTime)

        if nonceAge > staleDuration {
            // NONCE ESKİMİŞ (STALE)!
            // İstemciye yeni bir nonce ve "stale=true" direktifi ile yanıt ver.
            newNonce := generateSimpleNonce()
            nonceMap[newNonce] = time.Now() // Yeni nonce'ı da kaydet

            // Eski nonce'ı map'ten silebiliriz (isteğe bağlı)
            delete(nonceMap, clientNonce)

            algoForChallenge := algorithmParam
            if algoForChallenge == "" { algoForChallenge = "MD5" }

            staleResponseHeader := fmt.Sprintf(`Digest realm="%s", qop="%s", nonce="%s", algorithm="%s", stale=true`,
                realm, qopParam, newNonce, algoForChallenge)

            context.Header("WWW-Authenticate", staleResponseHeader)
            context.JSON(http.StatusUnauthorized, gin.H{"authenticated": false, "stale": true})
            return
        }
    }
    // --- STALE KONTROLÜ SONA ERDİ ---
    // Eğer kod buraya ulaştıysa, nonce tazedir ve geçerlidir.

    // Buradan sonrası, daha önce yazdığın imza doğrulama mantığının aynısı...
    clientUsername := authParams["username"]
    clientRealm := authParams["realm"]
    clientURI := authParams["uri"]
    clientQop := authParams["qop"]
    clientNC := authParams["nc"]
    clientCNonce := authParams["cnonce"]
    clientResponse := authParams["response"]

	if clientRealm != realm {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Realm mismatch"})
		return
	}
	// İstemcinin gönderdiği kullanıcı adı ile URL'deki kullanıcı adı aynı olmalı.
	if clientUsername != userParam {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Username mismatch"})
		return
	}

	ha1 := ""
	ha2 := ""
	expectedResponse := ""

	switch algorithmParam {
	case "MD5":
		ha1 = md5Hash(fmt.Sprintf("%s:%s:%s", userParam, clientRealm, passwdParam))
		ha2 = md5Hash(fmt.Sprintf("%s:%s", context.Request.Method, clientURI))
		expectedResponse = md5Hash(fmt.Sprintf("%s:%s:%s:%s:%s:%s", ha1, clientNonce, clientNC, clientCNonce, clientQop, ha2))
	case "SHA-256":
		ha1 = sha256Hash(fmt.Sprintf("%s:%s:%s", userParam, clientRealm, passwdParam))
		ha2 = sha256Hash(fmt.Sprintf("%s:%s", context.Request.Method, clientURI))
		expectedResponse = sha256Hash(fmt.Sprintf("%s:%s:%s:%s:%s:%s", ha1, clientNonce, clientNC, clientCNonce, clientQop, ha2))

	case "":
		ha1 = md5Hash(fmt.Sprintf("%s:%s:%s", userParam, clientRealm, passwdParam))
		ha2 = md5Hash(fmt.Sprintf("%s:%s", context.Request.Method, clientURI))
		expectedResponse = md5Hash(fmt.Sprintf("%s:%s:%s:%s:%s:%s", ha1, clientNonce, clientNC, clientCNonce, clientQop, ha2))
	default:
        // Desteklenmeyen bir algoritma istendi
		response["error"] = "Invalid digest response"
        context.JSON(http.StatusBadRequest, response)
        return // Fonksiyonu burada bitir.
    }
	
	if clientResponse == expectedResponse {
		response["authenticated"] = true
		response["user"] = userParam
		context.JSON(http.StatusOK, response)

	} else {
		response["authenticated"] = false
		response["error"] = "Invalid digest response"
		context.JSON(http.StatusUnauthorized, response)
	}

}

func generateSimpleNonce() string {
	// Şimdilik zaman damgasını nonce olarak kullanalım.
	// Bu güvenli değil ama mantığı anlamak için yeterli.
	return fmt.Sprintf("%x", time.Now().UnixNano())
}

// Authorization: Digest username="...", realm="...", ...
func parseDigestAuthHeader(header string) map[string]string {
	//header must start with Digest
	if !strings.HasPrefix(header, "Digest ") {
		return nil
	}
	//delete Digest part
	trimmed := strings.TrimPrefix(header, "Digest ")

	params := make(map[string]string)
	//virgülle ayrılmış kısımlara böl
	parts := strings.Split(trimmed, ",")

	for _, part := range parts {
		// Her kısmı "=" ile anahtar ve değere ayır
		keyValue := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(keyValue) == 2 {
			key := keyValue[0]
			// Değerler genellikle tırnak içinde gelir, bu tırnakları temizle
			value := strings.Trim(keyValue[1], "\"")
			params[key] = value
		}
	}
	return params
}

// verifies basic authentication but does not show prompt on the browser
func VerifyHiddenBasicAuth(context *gin.Context) {
	// header, err := getAuthorizationHeader(context)
	// if err != nil {
	// 	//don't add WWW-Authenticate header to preevent propmt from showing up
	// 	context.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// // Basic validation
	// if !isValidBasicAuth(header) {
	// 	context.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Invalid username or password.",
	// 	})
	// 	return
	// }

	// context.JSON(http.StatusOK, gin.H{
	// 	"message": "Authentication successful.",
	// })
}

func getAuthorizationHeader(context *gin.Context) (string, error) {
	header := context.GetHeader("Authorization")

	if header == "" {
		return "", errors.New("Authorization header is empty. Please enter valid credentials")

	}
	return header, nil
}

// farklı girdilerle aynı md5 sum üretilebiliyor. bunun için sonradan bundan vazgeçildi
func md5Hash(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func sha256Hash(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
