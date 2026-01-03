package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// BearerAuthHandler godoc
// @Summary      Bearer Authentication
// @Description  Authenticate using Bearer token
// @Tags         auth
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Success      200 {object} map[string]interface{}
// @Failure      401 {object} map[string]string
// @Router       /bearer [get]
func BearerAuthHandler(c *gin.Context) {
	//Header var mı? → yoksa 401
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// “Bearer ” ile başlıyor mu? → hayırsa 401
	if !strings.HasPrefix(authorizationHeader, "Bearer ") {
		c.JSON(401, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	// Token’ı al → doğruysa 200, yanlışsa yine 401
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	if token == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(200, gin.H{
		"authenticated": true,
	})
}
