package auth

import (
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuthHandler(c *gin.Context) {
	userParam := c.Param("user")
	passParam := c.Param("passwd")
	authorizationHeader := c.GetHeader("Authorization")

	if authorizationHeader == "" {
		c.Header("WWW-Authenticate", `Basic realm="Fake Realm"`)
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if !strings.HasPrefix(authorizationHeader, "Basic ") {
		c.JSON(400, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	encodedPart := strings.TrimPrefix(authorizationHeader, "Basic ")
	decodedAuth, err := base64Decode(encodedPart)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid base64 encoding"})
		return
	}

	if decodedAuth != userParam+":"+passParam {
		c.Header("WWW-Authenticate", `Basic realm="Fake Realm"`)
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(200, gin.H{
		"authenticated": true,
		"user":          userParam,
	})
}

func base64Decode(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
