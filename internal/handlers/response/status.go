package response

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StatusHandler(c *gin.Context) {
	codeStr := c.Param("code")
	code ,err := strconv.Atoi(codeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if code < 100 || code > 599 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status code must be between 100 and 599"})
		return
	}
	if code == http.StatusTeapot {
		c.String(http.StatusTeapot, "I'm a teapot")
		return
	}
	c.Status(code)
}