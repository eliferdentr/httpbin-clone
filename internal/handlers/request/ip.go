package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
IP Handler (router /ip)

Döndürmesi gereken JSON:

{
  "origin": "X.X.X.X"
}

C.ClientIP() kullan.
Proxy header'ları otomatik işler.
*/

// IPHandler godoc
//
// @Summary      Get client IP
// @Description Returns origin IP address
// @Tags         request
// @Produce      application/json
// @Success      200 {object} map[string]string
// @Router       /ip [get]
func IPHandler(c *gin.Context) {
	// 1) client ip al
	// ip := c.ClientIP()
	ip := c.ClientIP()
	// 2) JSON döndür
	c.JSON(http.StatusOK, gin.H{
		"origin": ip,
	})
}
