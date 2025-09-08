package response

import (
	"crypto/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils"
)

func BytesHandler(c *gin.Context) {
	n, invalid := utils.GetNParam(c)
	if invalid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bytes count"})
		return
	}

	const maxBytes = 100 * 1024

	if n > maxBytes {
		c.JSON(http.StatusBadRequest, gin.H{"error": "too many bytes requested"})
		return
	}
	data := make([]byte, n)
	if _, err := rand.Read(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate random bytes"})
		return
	}

	// Binary response
	c.Data(http.StatusOK, "application/octet-stream", data)

}
