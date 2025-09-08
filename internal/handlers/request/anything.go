package request

import (
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils"
)

func AnythingHandler(c *gin.Context) {
	args := utils.GetKeyValueMap(c.Request.URL.Query())
	headers := utils.GetKeyValueMap(c.Request.Header)
	form := utils.GetKeyValueMap(c.Request.PostForm)
	method := c.Request.Method
	url := c.Request.URL.String()
	origin := c.ClientIP()

	rawBody, jsonBody := utils.GetJSONBody(c)

	c.JSON(200, gin.H{
		"args":    args,
		"data":    rawBody,   // ham body string
		"form":    form,      // form-encoded
		"json":    jsonBody,  // parse edildiyse JSON, yoksa nil
		"headers": headers,
		"method":  method,
		"origin":  origin,
		"url":     url,
	})
}