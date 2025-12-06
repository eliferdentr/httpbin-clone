package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HeadersHandler(c *gin.Context) {
	// 1) header'ları al
	headers := c.Request.Header
	// 2) JSON formatında döndür
	// {
	//   "headers": { ... }
	// }
	headersMap := map[string]interface{}{}
	for k, v := range headers {
		headersMap[k] = v
	}

	c.JSON(http.StatusOK, gin.H{
		"headers": headersMap,
	})
}
