package images

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PNGHandler(c *gin.Context) {
	data, err := os.ReadFile("testdata/image.png")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "png not found"})
		return
	}

	c.Data(http.StatusOK, "image/png", data)
}
