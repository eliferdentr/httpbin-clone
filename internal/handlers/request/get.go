package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils"
)

func GetHandler(c *gin.Context) {
	args := utils.GetKeyValueMap(c.Request.URL.Query())
	headers := utils.GetKeyValueMap(c.Request.Header)

	c.JSON(http.StatusOK, gin.H{
		"args":    args,
		"headers": headers,
		"url":     c.Request.URL.String(),
	})

}
