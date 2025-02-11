package controllers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"go/constant"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	utils "httbinclone-eliferden.com/utils"
	constants "httbinclone-eliferden.com/utils/constants"
)

//verifies basic authentication and returns 200 if successful
func VerifyBasicAuth(context *gin.Context) {
	header, err := getAuthorizationHeader(context)
	if err != nil {
		//trigger prompt by adding WWW-Authenticate
		context.Header("WWW-Authenticate", `Basic realm="Access to the site"`)
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Basic auth validation
	if !isValidBasicAuth(header) {
		context.Header("WWW-Authenticate", `Basic realm="Access to the site"`)
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Authentication successful.",
	})
}


//verifies basic authentication but does not show prompt on the browser
func VerifyHiddenBasicAuth(context *gin.Context) {
	header, err := getAuthorizationHeader(context)
	if err != nil {
		//don't add WWW-Authenticate header to preevent propmt from showing up
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Basic validation
	if !isValidBasicAuth(header) {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Authentication successful.",
	})
}


//makes HTTP Digest Authentication
func VerifyDigestAuth(context *gin.Context) {

	header,err := getAuthorizationHeader(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error" : err})
	}
	if header == "" {
		//create nonce
	nonce, err := utils.GenerateNonce(constants.NONCE_BYTE_LENGTH)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error" : "An error occured while generating the nonce"})
	}

	digestString := fmt.Sprintf(`Digest realm="Access to the site", nonce="%d", algorithm=%s`, nonce, constants.NONCE_HASHING_ALGORITHM)
	context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing."})
	context.Header("WWW-Authenticate", digestString)
	return
	}

		// 3. Authorization header'ını parçala ve bilgileri al (username, nonce, response, vb.)
		authInfo := parseDigestAuthHeader(header)




}

func getAuthorizationHeader (context *gin.Context) (string, error) {
	header := context.Request.Header.Get("Authorization")

	if  header == "" {
		return "", errors.New ("authorization header is empty. Please enter valid credentials")
	}
	return header, nil
}

func getAuthorizationParts (header string) error {
	authParts := strings.SplitN(header," ", 2)
	if len(authParts) != 2 || authParts[0] != "Basic" {
		return errors.New("invalid Authorization header format. Expected 'Basic'")
	}
	payload, _ := base64.StdEncoding.DecodeString(authParts[1])
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 || pair[0] != "username" || pair[1] != "password" {
		return errors.New ("Invalid username or password")
	}
	return nil
}

func isValidBasicAuth(header string) bool {
	authParts := strings.SplitN(header, " ", 2)
	if len(authParts) != 2 || authParts[0] != "Basic" {
		return false
	}

	payload, _ := base64.StdEncoding.DecodeString(authParts[1])
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 || pair[0] != "username" || pair[1] != "password" {
		return false
	}

	return true
}

func parseDigestAuthHeader(header string) map[string]string{
  //remove digest prefix
  if !strings.HasPrefix(header, "Digest ") {
	return nil
  }

  header = strings.TrimPrefix(header, "Digest ")

  authMap := make(map[string]string)

  pairs := utils.SplitByCommas(header)

  for _, pair := range pairs {
	key, value := utils.ExtractKeyValue (pair)
	authMap[key] = value
  }

  return authMap 


	
}

