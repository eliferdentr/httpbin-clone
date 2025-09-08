package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils"
)

func StreamHandler(c *gin.Context) {
	n, invalid := utils.GetNParam(c)
	if invalid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid stream count"})
		return
	}
	c.Header("Content-Type", "application/json")

	for i := 0; i < n; i++ {
		// /get tarzı response hazırla
		item := gin.H{
			"args":    utils.GetKeyValueMap(c.Request.URL.Query()),
			"headers": utils.GetKeyValueMap(c.Request.Header),
			"origin":  c.ClientIP(),
			"url":     c.Request.URL.String(),
		}

		// JSON'a çevir
		b, err := json.Marshal(item)
		if err != nil {
			return
		}

		// Yaz + newline
		c.Writer.Write(b)
		c.Writer.Write([]byte("\n"))
		c.Writer.Flush()

		// İstersen biraz beklet
		// time.Sleep(200 * time.Millisecond)
	}
}
