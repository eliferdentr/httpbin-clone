package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	utils "httbinclone-eliferden.com/utils"
	constants "httbinclone-eliferden.com/utils/constants"
)

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

// makes HTTP Digest Authentication
func VerifyDigestAuth(context *gin.Context) {
	response := utils.BuildResponse(context, nil)
	header, err := getAuthorizationHeader(context)
	//get path variables
	qopParam := context.Param("qop")
	userParam := context.Param("user")
	passwdParam := context.Param("passwd")

	if err != nil {
		//challenge phase
		//generate a realm
		realm := fmt.Sprintf("%s@eliferden.com", userParam)

		//generate nonce and opaque

		response["error"] = err.Error()
		context.JSON(http.StatusUnauthorized, response)
		return
	}





	// // Eğer Authorization header yoksa, nonce oluştur ve istemciye gönder
	// if header == "" {
	// 	nonce, err := utils.GenerateNonce(constants.NONCE_BYTE_LENGTH)
	// 	if err != nil {
	// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while generating the nonce"})
	// 		return
	// 	}

	// 	digestString := fmt.Sprintf(`Digest realm="Access to the site", nonce="%s", algorithm=%s`, nonce, constants.NONCE_HASHING_ALGORITHM)
	// 	context.Header("WWW-Authenticate", digestString)
	// 	context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing."})
	// 	return
	// }

	// // Authorization header'ını parçala ve bilgileri al
	// authInfo := parseDigestAuthHeader(header)
	// if authInfo == nil {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format."})
	// 	return
	// }

	// // Nonce kontrolü
	// if authInfo["nonce"] == "" || authInfo["nonce"] != constants.EXPECTED_NONCE {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid nonce value."})
	// 	return
	// }

	// // Kullanıcı adı ve şifre doğrulama
	// if authInfo["username"] != constants.EXPECTED_USERNAME || authInfo["response"] != calculateDigestHash(authInfo) {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password."})
	// 	return
	// }

	// context.JSON(http.StatusOK, gin.H{"message": "Digest authentication successful."})
}

func parseDigestAuthHeader(header string) map[string]string {
	if !strings.HasPrefix(header, "Digest ") {
		return nil
	}
	header = strings.TrimPrefix(header, "Digest ")

	authMap := make(map[string]string)
	pairs := utils.SplitByCommas(header)

	for _, pair := range pairs {
		key, value := utils.ExtractKeyValue(pair)
		authMap[key] = value
	}

	return authMap
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

func VerifyDigestAuthWithAlgortihm(context *gin.Context) {
	header, err := getAuthorizationHeader(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Eğer Authorization header yoksa, nonce oluştur ve istemciye gönder
	if header == "" {
		nonce, err := utils.GenerateNonce(constants.NONCE_BYTE_LENGTH)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while generating the nonce"})
			return
		}

		digestString := fmt.Sprintf(`Digest realm="Access to the site", nonce="%s", algorithm=%s`, nonce, constants.NONCE_HASHING_ALGORITHM)
		context.Header("WWW-Authenticate", digestString)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing."})
		return
	}

	// Authorization header'ını parçala ve bilgileri al
	authInfo := parseDigestAuthHeader(header)
	if authInfo == nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format."})
		return
	}

	// Nonce kontrolü
	if authInfo["nonce"] == "" || authInfo["nonce"] != constants.EXPECTED_NONCE {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid nonce value."})
		return
	}

	// Kullanıcı adı ve şifre doğrulama
	if authInfo["username"] != constants.EXPECTED_USERNAME || authInfo["response"] != calculateDigestHash(authInfo) {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Digest authentication successful."})
}

func getAuthorizationHeader(context *gin.Context) (string, error) {
	header := context.GetHeader("Authorization")

	if header == "" {
		return "", errors.New("Authorization header is empty. Please enter valid credentials")

	}
	return header, nil
}



func calculateDigestHash(authInfo map[string]string) string {
	ha1 := md5Hash(authInfo["username"] + ":" + constants.REALM + ":" + constants.EXPECTED_PASSWORD)
	ha2 := md5Hash(authInfo["method"] + ":" + authInfo["uri"])
	return md5Hash(ha1 + ":" + authInfo["nonce"] + ":" + ha2)
}

func md5Hash(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
