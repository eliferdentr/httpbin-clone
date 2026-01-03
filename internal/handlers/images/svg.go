package images

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// SVGHandler godoc
//
// @Summary      Get SVG image
// @Description  Returns a sample SVG image
// @Tags         images
// @Produce      image/svg+xml
// @Success      200 {string} string "SVG XML"
// @Failure      500 {object} map[string]string
func SVGHandler(c *gin.Context) {
	data, err := os.ReadFile("testdata/image.svg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "svg not found"})
		return
	}

	c.Data(http.StatusOK, "image/svg+xml", data)
}
