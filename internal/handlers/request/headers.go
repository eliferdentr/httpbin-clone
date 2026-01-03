package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HeadersHandler godoc
//
// @Summary      Get request headers
// @Description Returns all request headers
// @Tags         request
// @Produce      application/json
// @Success      200 {object} map[string]interface{}
// @Router       /headers [get]
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
