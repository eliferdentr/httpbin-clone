package response

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils"
)

func RangeHandler(c *gin.Context) {
  n, invalid := utils.GetNParam(c)
  if invalid {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid range count"})
    return
  }
// Default değerler
	duration, _ := strconv.Atoi(c.DefaultQuery("duration", "0"))
	chunkSize, _ := strconv.Atoi(c.DefaultQuery("chunk_size", strconv.Itoa(n)))

	// Data hazırla: 0..n-1 byte
	data := make([]byte, n)
	for i := 0; i < n; i++ {
		data[i] = byte(i % 256)
	}

	c.Header("Content-Type", "application/octet-stream")
// Streaming (duration verilmişse)
	if duration > 0 && chunkSize > 0 && chunkSize < n {
		interval := float64(duration) / float64(n/chunkSize)
		c.Status(http.StatusOK)
		c.Writer.WriteHeaderNow()

		for i := 0; i < n; i += chunkSize {
			end := i + chunkSize
			if end > n {
				end = n
			}
			_, err := c.Writer.Write(data[i:end])
			if err != nil {
				return
			}
			c.Writer.Flush()
			time.Sleep(time.Duration(interval * float64(time.Second)))
		}
	} else {
		// Tek seferde gönder
		c.Data(http.StatusOK, "application/octet-stream", data)
	}
  
}