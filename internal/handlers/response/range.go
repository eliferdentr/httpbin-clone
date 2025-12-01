package response

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/*
HttpBin’de /range/:n şöyle çalışır:

1. n kadar byte üretir
Ama bytes handler’daki gibi random değil:
Hepsi 0x00, 0x01, 0x02... şeklinde artan değerler.
Örneğin /range/5 output:
\x00\x01\x02\x03\x04
Tabii binary olduğu için görünmez ama tam n byte döner.

2. Chunked transfer encoding yapılır
Kod örnekleri gibi:
chunk’lar 1 byte olacak
her byte arası mini sleep yok (drip’ten farklı)
ama flush yapılmalı

3. Query parametresi olarak duration olabilir
HttpBin’de opsiyonel:
/range/10?duration=2
→ 10 byte, 2 saniye içinde akar
→ chunkInterval = duration / numbytes

4. Header'lar:
Content-Type: application/octet-stream

5. n < 0 veya sayı değil → 400
*/
func RangeHandler(c *gin.Context) {
	nStr := c.Param("n")
	if nStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "n parameter cannot be empty"})
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil || n < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid n parameter"})
		return
	}

	// duration param (optional)
	duration := 0
	durationStr := c.Query("duration")
	if durationStr != "" {
		duration, err = strconv.Atoi(durationStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if duration < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "duration cannot be negative"})
			return
		}
	}

	// headers
	c.Status(http.StatusOK)
	c.Header("Content-Type", "application/octet-stream")

	writer := c.Writer

	// interval per byte
	var interval time.Duration
	if duration > 0 && n > 0 {
		interval = time.Duration(duration) * time.Second / time.Duration(n)
	} else {
		interval = 0
	}

	// write bytes
	for i := 0; i < n; i++ {
		writer.Write([]byte{byte(i % 256)})
		writer.Flush()
		time.Sleep(interval)
	}
}

