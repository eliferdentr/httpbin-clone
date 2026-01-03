package response

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
HttpBin’de /stream/:n şöyle çalışıyor:

1) n adet JSON satırı gönderir
Her satır ayrı bir chunk olarak akar.
Örnek /stream/3 çıktısı:
{"id":0}
{"id":1}
{"id":2}

2) Content-Type: application/json
Ama chunked transfer olduğu için body bir bütün JSON değil — satır satır akar.

3) Header'lar:
Content-Type: application/json

4) n negatif veya sayı değil → 400

5) Her chunk yazıldıktan sonra Flush() çağrılır.
*/

// StreamHandler godoc
//
// @Summary      Stream JSON lines
// @Description Streams N JSON objects line by line
// @Tags         response
// @Param        n path int true "Number of lines"
// @Produce      application/json
// @Success      200 {string} string
// @Failure      400 {object} map[string]string
// @Router       /stream/{n} [get]
func StreamHandler(c *gin.Context) {
	// 1) n param al → parse et → negatifse 400
	nStr := c.Param("n")
	if nStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "n parameter cannot be empty"})
		return
	}
	n, err := strconv.Atoi(nStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if n < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "n parameter cannot be negative"})
	}

	// 2) Content-Type ayarla
	c.Header("Content-Type", "application/json")

	// 3) Status 200
	c.Status(http.StatusOK)

	writer := c.Writer

	for i := 0; i < n; i++ {
		line := fmt.Sprintf(`{"id": %d}`+"\n", i)
		writer.Write([]byte(line))
		writer.Flush()
	}

}
