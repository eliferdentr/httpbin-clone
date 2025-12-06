package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
IP Handler (httpbin /ip)

Döndürmesi gereken JSON:

{
  "origin": "X.X.X.X"
}

C.ClientIP() kullan.
Proxy header'ları otomatik işler.
*/

func IPHandler(c *gin.Context) {
	// 1) client ip al
	// ip := c.ClientIP()
	ip := c.ClientIP()
	// 2) JSON döndür
	c.JSON(http.StatusOK, gin.H{
		"origin": ip,
	})
}
