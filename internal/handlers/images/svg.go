package images

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SVGHandler(c *gin.Context) {
	data, err := os.ReadFile("testdata/image.svg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "svg not found"})
		return
	}

	c.Data(http.StatusOK, "image/svg+xml", data)
}
