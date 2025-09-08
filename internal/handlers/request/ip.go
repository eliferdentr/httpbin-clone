package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IPHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ip": c.ClientIP()})
}