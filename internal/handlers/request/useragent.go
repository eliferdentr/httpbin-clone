package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User-Agent = isteği hangi program / cihaz gönderiyor

// UserAgentHandler godoc
//
// @Summary      Get User-Agent
// @Description Returns client User-Agent string
// @Tags         request
// @Produce      application/json
// @Success      200 {object} map[string]string
// @Router       /user-agent [get]
func UserAgentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user-agent": c.Request.UserAgent()})
}
