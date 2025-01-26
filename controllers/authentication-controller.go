package controllers

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	utils "httbinclone-eliferden.com/utils"
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
	// nonceService, err  := utils.GenerateNonce()
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{
	// 		"error" : "There was an error while generating digest auth",
	// 	})
	// }




}

func getAuthorizationHeader (context *gin.Context) (string, error) {
	header := context.Request.Header.Get("Authorization")

	if  header == "" {
		return "", errors.New ("Authorization header is empty. Please enter valid credentials.")
	}
	return header, nil
}

func getAuthorizationParts (header string) error {
	authParts := strings.SplitN(header," ", 2)
	if len(authParts) != 2 || authParts[0] != "Basic" {
		return errors.New("Invalid Authorization header format. Expected 'Basic'.")
	}
	payload, _ := base64.StdEncoding.DecodeString(authParts[1])
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 || pair[0] != "username" || pair[1] != "password" {
		return errors.New ("Invalid username or password.")
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

