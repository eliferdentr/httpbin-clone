package images

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// JpegHandler godoc
//
// @Summary      Get JPEG image
// @Description  Returns a sample JPEG image
// @Tags         images
// @Produce      image/jpeg
// @Success      200 {file} jpeg
// @Failure      500 {object} map[string]string
// @Router       /image/jpeg [get]
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
