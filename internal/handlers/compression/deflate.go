package compression

import (
	"compress/zlib"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
DeflateHandler

Davranış:
  - Accept-Encoding header içinde "deflate" varsa:
    → response deflate (zlib) ile sıkıştırılır
  - Yoksa:
    → normal JSON döner

JSON formatı:

	{
	  "deflated": true / false,
	  "headers": {...},
	  "method": "GET"
	}

Deflate aktifse:
- Content-Encoding: deflate
- Body zlib writer ile yazılır
*/

// DeflateHandler godoc
//
// @Summary      Deflate compressed response
// @Description  Returns a JSON response, deflate-compressed if client supports it (Accept-Encoding: deflate)
// @Tags         compression
// @Produce      application/json
// @Param        Accept-Encoding header string false "deflate"
// @Success      200 {object} map[string]interface{}
// @Failure      500 {object} map[string]string
// @Router       /deflate [get]
func DeflateHandler(c *gin.Context) {

	isDeflate := strings.Contains(
		strings.ToLower(c.GetHeader("Accept-Encoding")),
		"deflate",
	)

	resp := map[string]interface{}{
		"deflated": isDeflate,
		"headers":  c.Request.Header,
		"method":   c.Request.Method,
	}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/json")

	// Deflate YOKSA → normal JSON
	if !isDeflate {
		c.JSON(http.StatusOK, resp)
		return
	}

	// Deflate VARSA
	c.Header("Content-Encoding", "deflate")
	c.Status(http.StatusOK)

	writer := zlib.NewWriter(c.Writer)
	defer func() {
		if err := writer.Close(); err != nil {
			log.Println("deflate close error:", err)
		}
	}()

	if _, err := writer.Write(jsonBytes); err != nil {
		log.Println("deflate write error:", err)
	}
}
