package images

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func JpegHandler(c *gin.Context) {
	// 1) JPEG dosyasını oku (örn: testdata/sample.jpg)
	//    - os.ReadFile kullan
	//    - hata varsa 500 dön
	// 1) JPEG dosyasını raw olarak oku
	//TODO : path düzenle
	jpegBytes, err := os.ReadFile("testdata/uia.jpg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not read jpeg file",
		})
		return
	}
	// 2) Header:
	//    Content-Type: image/jpeg
	c.Header("Content-Type", "image/jpeg")

	// 3) c.Data ile response yaz
	//    status = 200
	//    body = jpeg bytes
	c.Data(http.StatusOK, "image/jpeg", jpegBytes)

}
