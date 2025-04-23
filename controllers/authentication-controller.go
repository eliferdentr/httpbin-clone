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
	payloadParts := getBasicAuthParts(header)

	if len(payloadParts) != 2 {
		response["error"] = "Failed to parse Basic auth header"
		context.Header("WWW-Authenticate", `Basic realm="Access to the site"`)
		context.JSON(http.StatusUnauthorized, response)
		return
	}

	userParam := context.Param("user")
    passParam := context.Param("password")
	
	if payloadParts [0] != userParam || payloadParts[1] != passParam {
		response["error"] = "Credentials don't match URL params"
		context.Header("WWW-Authenticate", `Basic realm="Access to the site"`)
		context.JSON(http.StatusUnauthorized, response)
		return
	}

	response["message"] = "Authentication successful"
	context.JSON(http.StatusOK, response)
}

func VerifyBearerAuth(context *gin.Context) {

	header, err := getAuthorizationHeader(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error" : err.Error(),
		})
		return
	}

	//split as "Bearer TOKEN"
	parts := strings.SplitN(header, " ", 2)

	if len(parts) != 2 || parts[0] != "Bearer" {
		context.JSON(http.StatusUnauthorized, 
			gin.H{"error": "Authorization must be Bearer {token}"})
        return

	}

	token := parts[1]
	if token != constants.EXPECTED_NONCE {
		context.JSON(
			http.StatusForbidden, gin.H{"error": "Invalid or expired token"})
        return
	}

	response := utils.BuildResponse(context, nil)
	response["message"] = "Bearer authentication successful"
	context.JSON(http.StatusOK, response)

}

//makes HTTP Digest Authentication
func VerifyDigestAuth(context *gin.Context) {
	response := utils.BuildResponse(context, nil)
	header, err := getAuthorizationHeader(context)
	if err != nil {
		response["error"] = err.Error()
		context.JSON(http.StatusInternalServerError, response)
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


//verifies basic authentication but does not show prompt on the browser
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





func getAuthorizationHeader (context *gin.Context) (string, error) {
	//expected: Authorization: Basic <base64(username:password)>
	header := context.Request.Header.Get("Authorization")

	if  header == "" {
		return "", errors.New("authorization header is empty. Please enter valid credentials")

	}
	return header, nil
}



func getBasicAuthParts (header string) [] string  {
	authParts := strings.SplitN(header, " ", 2)
	if len(authParts) != 2 || authParts[0] != "Basic" {
		return []string{}
	}

	payload, _ := base64.StdEncoding.DecodeString(authParts[1])
	parts := strings.SplitN(string(payload), ":", 2)
	if len(parts) != 2 {
		return []string{}
	}
	return parts
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

