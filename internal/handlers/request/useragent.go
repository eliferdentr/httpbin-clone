package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//User-Agent = isteği hangi program / cihaz gönderiyor
func UserAgentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user-agent": c.Request.UserAgent()})
}
