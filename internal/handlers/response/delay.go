package response

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/internal/handlers/request"
)

func DelayHandler(c *gin.Context) {
	secondsParam := c.Param("seconds")
	seconds, err := strconv.Atoi(secondsParam)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid seconds parameter"})
		return
	}

	var sleepTime time.Duration
	if seconds < 0 {
		seconds = 0
	}
	if seconds > 10 {
		seconds = 10
	}
	time.Sleep(time.Duration(seconds) * time.Second)
	time.Sleep(sleepTime)
	request.GetHandler(c)
}
