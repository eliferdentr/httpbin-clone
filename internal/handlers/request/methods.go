package request

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
MethodsHandler
- Sadece gelen HTTP method'unu döndürür
- Body varsa body
- Query param varsa query
*/

func MethodsHandler(c *gin.Context) {
	method := c.Request.Method

	headers := c.Request.Header

	query := c.Request.URL.Query()

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"headers": headers,
		"query":   query,
		"body":    string(body),
		"method":  method,
	})

}
