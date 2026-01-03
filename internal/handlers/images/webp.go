package images

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// WEBPHandler godoc
//
// @Summary      Get WEBP image
// @Description  Returns a sample WEBP image
// @Tags         images
// @Produce      image/webp
// @Success      200 {file} webp
// @Failure      500 {object} map[string]string
// @Router       /image/webp [get]
func WEBPHandler(c *gin.Context) {
	data, err := os.ReadFile("testdata/image.webp")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "webp not found"})
		return
	}

	c.Data(http.StatusOK, "image/webp", data)
}
