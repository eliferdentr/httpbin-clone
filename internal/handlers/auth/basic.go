package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils"
)

// BasicAuthHandler godoc
//
// @Summary      Basic Authentication
// @Description  Basic Auth with username and password
// @Tags         auth
// @Produce      json
// @Param        user   path  string  true  "Username"
// @Param        passwd path  string  true  "Password"
// @Success      200 {object} map[string]interface{}
// @Failure      401 {object} map[string]string
// @Router       /basic-auth/{user}/{passwd} [get]
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
	decodedAuth, err := utils.Base64Decode(encodedPart)
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
