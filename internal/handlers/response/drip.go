package response

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func DripHandler(c *gin.Context) {
	numbytes, _ := strconv.Atoi(c.DefaultQuery("numbytes", "10"))
	duration, _ := strconv.Atoi(c.DefaultQuery("duration", "2"))
	delay, _ := strconv.Atoi(c.DefaultQuery("delay", "0"))
	code, _ := strconv.Atoi(c.DefaultQuery("code", "200"))

	if code < 100 || code > 599 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status code must be between 100 and 599"})
		return
	}
	if numbytes <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "numbytes must be > 0"})
		return
	}
	if delay > 0 {
		time.Sleep(time.Duration(delay) * time.Second)
	}

	interval := float64(duration) / float64(numbytes)

	c.Status(code)
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.WriteHeaderNow()

	for i := 0; i < numbytes; i++ {
		_, err := c.Writer.Write([]byte("*")) 
		if err != nil {
			return
		}
		c.Writer.Flush()
		time.Sleep(time.Duration(interval * float64(time.Second)))
	}

}
