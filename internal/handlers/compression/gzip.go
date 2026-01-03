package compression

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
GzipHandler
  - Accept-Encoding header içinde "gzip" varsa → gzip sıkıştırılmış body döner
  - Yoksa → normal JSON döner
  - HttpBin gibi:
    {
    "gzipped": true/false,
    "headers": ...    // tüm header'lar
    "method": ...     // GET, POST
    }

Body gzip'liyse:
- Content-Encoding: gzip
- Body gzip writer ile yazılır
*/

// GzipHandler godoc
//
// @Summary      Gzip compressed response
// @Description Returns a JSON response, gzip-compressed if client supports it
// @Tags         compression
// @Produce      application/json
// @Param        Accept-Encoding header string false "gzip"
// @Success      200 {object} map[string]interface{}
// @Failure      500 {object} map[string]string
// @Router       /gzip [get]
func GzipHandler(c *gin.Context) {
	/* 1) Client’ın Accept-Encoding header’ına bakılır:
	- Eğer "gzip" içeriyorsa → client gzip sıkıştırmasını kabul ediyor demektir.
	- Eğer yoksa → normal (gzip'siz) JSON döndürülür.
	*/

	/* 2) Gzip aktifse:
	- Response header’a: Content-Encoding: gzip eklenir.
	- Response body gzip.Writer üzerinden sıkıştırılarak yazılır.
		Yani:
	gzipWriter := gzip.NewWriter(c.Writer)
	gzipWriter.Write(jsonBytes)
	gzipWriter.Close()

	Bu sayede client body’yi açmak zorundadır.
	3) Dönen JSON şu formatta olur:
		{
			"gzipped": true/false,
			"headers": {... request headers ...},
			"method": "GET" / "POST" ...
		}
	*/

	isGzip := strings.Contains(strings.ToLower(c.GetHeader("Accept-Encoding")), "gzip")

	jsonBodyMap := map[string]interface{}{}
	jsonBodyMap["gzipped"] = isGzip
	jsonBodyMap["headers"] = c.Request.Header
	jsonBodyMap["method"] = c.Request.Method
	jsonBody, err := json.Marshal(jsonBodyMap)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	/*
		5) gzip KAPALIYSA:
		- Normal c.JSON ile yazılır.
		- "gzipped": false döner.
		- Content-Encoding header **olmamalıdır**.

		6) Edge-Case’ler:
		- Accept-Encoding: gzip, compress, deflate → içinde gzip geçiyorsa yeterlidir.
		- Method fark etmez (GET, POST, PUT…).
		- Headers, method, gzipped flag her durumda JSON içinde döner.
			Bu handler HttpBin’in /gzip endpoint’inin birebir aynısıdır.
	*/
	c.Header("Content-Type", "application/json")
	if !isGzip {
		c.JSON(http.StatusOK, jsonBody)
		return
	}
	/*
		4) gzip ile yazıyorsan:
				- Kesinlikle c.JSON kullanamazsın.
					Çünkü c.JSON body’yi içerde encode edip direkt Writer'a basar.
				Gzip'ten geçmesi için senin manuel olarak json.Marshal yapıp, gzipWriter.Write ile yazman gerekir.
	*/
	c.Header("Content-Encoding", "gzip")
	c.Status(http.StatusOK)
	gzipWriter := gzip.NewWriter(c.Writer)
	defer func(gzipWriter *gzip.Writer) {
		err := gzipWriter.Close()
		if err != nil {
			log.Println("gzip close error:", err)
		}
	}(gzipWriter)
	_, err2 := gzipWriter.Write(jsonBody)
	if err2 != nil {
		log.Println("gzip write error:", err2)
		return
	}

}
