package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	services "httbinclone-eliferden.com/services/implementation"
)

//returns the list of all the endpoints
func RootHandler(context *gin.Context) {
	routes := gin.Default().Routes()
	if len(routes) < 1 {
		context.JSON(http.StatusOK, gin.H{
			"message" : "There were not any registered routes found.",
		})
		return
	}
	filteredRoutes := make([]string, len(routes))
	for _, route := range routes {
		filteredRoutes = append(filteredRoutes, route.Path)
	}
	context.JSON(http.StatusOK, gin.H{
		"message" : "There were not any registered routes found.",
	})
}

//returns a random uuid
func GetUUID(context *gin.Context) {
	uuidString := uuid.NewString()
	context.JSON(http.StatusOK, gin.H {
		"generated_uuid" : uuidString,
	})
}

//returns the USER-AGENT info of the request
func GetUserAgent(context *gin.Context) {
	userAgent := context.Request.Header.Get("User-Agent")
	context.JSON(http.StatusOK, gin.H{
		"user-agent": userAgent,
	})
}

//returns the headers
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
	context.JSON(http.StatusOK, gin.H{
		"headers": formattedHeaders,
	})
}

//returns the client ip
func GetIP(context *gin.Context) {
	clientIp := context.ClientIP()

	if clientIp == "" {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"error": "Client IP could not be resolved",
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
