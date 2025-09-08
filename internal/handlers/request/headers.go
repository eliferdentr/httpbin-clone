package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HeadersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, c.Request.Header)
}