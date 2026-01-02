package compression

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/gin-gonic/gin"
)

// BrotliHandler godoc
//
// @Summary      Brotli compressed response
// @Description  Returns a JSON response, brotli-compressed if client supports it (Accept-Encoding: br)
// @Tags         compression
// @Produce      application/json
// @Param        Accept-Encoding header string false "br"
// @Success      200 {object} map[string]interface{}
// @Failure      500 {object} map[string]string
// @Router       /brotli [get]
/*
BrotliHandler

Davranış:
- Accept-Encoding header içinde "br" varsa:
    → response brotli ile sıkıştırılır
- Yoksa:
    → normal JSON döner

JSON formatı (her iki durumda da):
{
  "brotli": true / false,
  "headers": {...},
  "method": "GET"
}

Brotli aktifse:
- Content-Encoding: br
- Body brotli writer ile yazılır
*/

func BrotliHandler(c *gin.Context) {

	// 1) Accept-Encoding header’ını oku
	//    → strings.Contains(..., "br") ile kontrol et
	isBrotli := strings.Contains(strings.ToLower(c.GetHeader("Accept-Encoding")), "br")

	jsonBodyMap := map[string]interface{}{}
	jsonBodyMap["brotli"] = isBrotli
	jsonBodyMap["headers"] = c.Request.Header
	jsonBodyMap["method"] = c.Request.Method
	jsonBody, err := json.Marshal(jsonBodyMap)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")

	// 4) Eğer brotli YOKSA:
	//    - c.JSON(200, obj)
	//    - Content-Encoding header EKLEME

	if !isBrotli {
		c.JSON(http.StatusOK, jsonBodyMap)
		return
	}

	// 5) Eğer brotli VARSA:
	//    - Header: Content-Encoding = br
	//    - Status = 200
	//    - brotli.NewWriter(c.Writer)
	//    - writer.Write(jsonBytes)
	//    - writer.Close()
	c.Header("Content-Encoding", "br")
	c.Status(http.StatusOK)
	bortliWriter := brotli.NewWriter(c.Writer)
	defer func(bortliWriter *brotli.Writer) {
		err := bortliWriter.Close()
		if err != nil {
			log.Println("brotli close error:", err)
		}
	}(bortliWriter)
	_, err2 := bortliWriter.Write(jsonBody)
	if err2 != nil {
		log.Println("brotli write error:", err2)
		return
	}
}
