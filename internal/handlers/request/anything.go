package request

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AnythingHandler(c *gin.Context) {
	// 1) METHOD
	method := c.Request.Method

	// 2) HEADERS
	headers := c.Request.Header

	// 3) QUERY PARAMS
	query := c.Request.URL.Query()

	// 4) FULL URL / PATH
	fullURL := c.Request.RequestURI

	// 5) BODY
	bodyBytes, _ := io.ReadAll(c.Request.Body)
	body := string(bodyBytes)

	// 6) JSON RESPONSE
	c.JSON(http.StatusOK, gin.H{
		"method":  method,
		"headers": headers,
		"query":   query,
		"url":     fullURL,
		"body":    body,
	})
}
