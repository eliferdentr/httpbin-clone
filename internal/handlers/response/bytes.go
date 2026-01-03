package response

import (
	"crypto/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BytesHandler godoc
//
// @Summary      Generate random bytes
// @Description Returns N random bytes
// @Tags         response
// @Param        n path int true "Number of bytes"
// @Produce      application/octet-stream
// @Success      200 {string} binary
// @Failure      400 {object} map[string]string
// @Router       /bytes/{n} [get]
func BytesHandler(c *gin.Context) {
	// 1) parametreyi al
	nStr := c.Param("n")
	if nStr == "" {
		c.JSON(400, gin.H{"error": "invalid n parameter"})
		return
	}
	// 2) integer'a çevir
	n, err := strconv.Atoi(nStr)
	// 3) hata varsa 400
	if err != nil || n < 0 {
		c.JSON(400, gin.H{"error": "invalid n parameter"})
		return
	}
	// 4) n < 0 ise 400
	// 5) make([]byte, n)
	nArr := make([]byte, n)
	// 6) rand.Read ile doldur
	_, err = rand.Read(nArr)
	if err != nil {
		c.JSON(500, gin.H{"error": "en internal error occured"})
		return
	}
	// 7) Content-Type: application/octet-stream
	c.Header("Content-Type", "application/octet-stream")
	// 8) c.Data(status=200, content-type, body)
	c.Data(http.StatusOK, "application/octet-stream", nArr)

}

/*
 1. n kadar random byte üretir

Binary data = []byte{ … }

 2. Content-Type → application/octet-stream

Çünkü bu raw byte.

 3. Body tam olarak n byte olmalı

Ne eksik ne fazla.

 4. Eğer n negatif veya parse edilemiyorsa → 400
 5. JSON değil, metin değil

Direkt binary response.
*/
