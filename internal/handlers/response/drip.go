package response

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/*
delay süresi bekle
status code set et
Response’u chunk chunk yaz (her chunk 1 byte)
Her byte arasında:
duration / numbytes kadar bekle
Chunked Transfer Encoding ile akar
(Gin bunu otomatik yapar Flush() ile)
Body şuna benzer (10 byte örneği):
XXXXXXXXXX
*/

func DripHandler(c *gin.Context) {
	// 1) query parametrelerini al:
	// numbytes
	// duration
	// delay
	// code (opsiyonel)

	code := http.StatusOK

	// 2) parse et → yanlışsa 400
	numbytes, err := strconv.Atoi(c.Query("numbytes"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if numbytes < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "numbytes must be non-negative"})
		return
	}
	duration, err := strconv.Atoi(c.Query("duration"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if duration < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Duration must be greater than 0"})
		return
	}
	delay, err := strconv.Atoi(c.Query("delay"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if delay < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delay must be greater than 0"})
		return
	}
	if c.Query("code") != "" {
		code, err = strconv.Atoi(c.Query("code"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
			if code < 100 || code > 599 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status code must be between 100 and 599"})
		return
	}
	}

	// 3) initial delay: time.Sleep(delay)
	time.Sleep(time.Duration(delay) * time.Second)
	// 4) status code set et: c.Status(code)
	c.Status(code)
	// 5) chunk süresi hesapla:
	//    chunkInterval := duration / numbytes
	chunkInterval := time.Duration(duration) * time.Second / time.Duration(numbytes)

	// 6) Writer'ı al
	//    w := c.Writer
	writer := c.Writer

	// 7) Döngü: numbytes kadar
	//       w.Write([]byte{byte(0x58)})  // 'X' veya random byte
	//       w.Flush()
	//       time.Sleep(chunkInterval)

	for i := numbytes; i > 0; i-- {
		writer.Write([]byte{byte(0x58)})
		writer.Flush()
		time.Sleep(chunkInterval)
	}
	// NOT: c.JSON vs yok → plain stream!

}
