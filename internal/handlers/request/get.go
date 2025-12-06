package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	// 1) Query parametreleri al
	// args := ...
	args := c.Request.URL.Query()

	// 2) Header'ları al
	// headers := ...
	headers := c.Request.Header

	// 3) origin (IP adresi) al
	// origin := ...
	ipAdress := c.ClientIP()

	// 4) full URL al
	// url := ...
	fullURL := c.Request.RequestURI

	// 5) JSON formatında döndür
	c.JSON(http.StatusOK, gin.H{
		"args":    args,
		"headers": headers,
		"origin":  ipAdress,
		"url":     fullURL,
	})
}
