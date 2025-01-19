package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "httbinclone-eliferden.com/services/implementation"
)

func GetHeaders(context *gin.Context) {
	headers := context.Request.Header
	if len(headers) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "No headers provided in the request",
		})
		return
	}
	service := services.NewRequestProcessorServiceImpl()
	formattedHeaders := service.GetHeader(headers)
	context.JSON(http.StatusOK, gin.H {
		"headers" : formattedHeaders,
	})
}

func GetIp(context *gin.Context) {
	clientIp := context.ClientIP()

	if clientIp == "" {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"error" : "Client IP could not be resolved",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"client_ip": clientIp,
		"meta": gin.H{
			"note": "This is the IP address detected by the server.",
		},
	})

}