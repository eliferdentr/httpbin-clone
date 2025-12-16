package images

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func WEBPHandler(c *gin.Context) {
	data, err := os.ReadFile("testdata/image.webp")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "webp not found"})
		return
	}

	c.Data(http.StatusOK, "image/webp", data)
}
