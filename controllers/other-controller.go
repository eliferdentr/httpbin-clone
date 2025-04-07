package controllers

import (
	"crypto/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	services "httbinclone-eliferden.com/services/implementation"
	"httbinclone-eliferden.com/utils"
	"httbinclone-eliferden.com/utils/constants"
)

//generates and returns random data of incoming bytes
func GetRange(context *gin.Context){
	n, err := utils.GetPositiveIntParam(context, "n")
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	data := make([]byte, n)
	_, err = rand.Read(data)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate random data.",
		})
		return
	}

	// Veriyi application/octet-stream tipinde döndür
	context.Data(http.StatusOK, "application/octet-stream", data)
}

//return an html page
func GetHTML( context *gin.Context){
	htmlContent := constants.HTMLCONTENT
	if htmlContent == "" {
		context.JSON(http.StatusInternalServerError,
			gin.H{"error": `There was an error while handling the request.
			HTML Template could not be found on the server.`})
			return
	}
	context.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
}

//return scanning rules of the site
func GetRobots ( context *gin.Context) {
	robotsContent := constants.ROBOTSCONTENT
	if robotsContent == "" {
		context.JSON(http.StatusInternalServerError,
			gin.H{"error": `There was an error while handling the request.
			robots.txt could not be found on the server.`})
			return
	}
	context.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(robotsContent))
}

//stops the incoming request and denies it
func GetDeny (context *gin.Context) {
context.AbortWithStatusJSON(http.StatusForbidden, gin.H{
	"error" : "Access Denied",
})
}

func GetImage (context *gin.Context) {
	
}

func GetAnything(context *gin.Context) {

	headers := context.Request.Header
	method := context.Request.Method
	params := context.Request.URL.Query()

	var body map[string]any
	if err := context.ShouldBindJSON(&body); err != nil {
		body = nil
	}

	service := services.NewRequestProcessorServiceImpl()
	details := service.GetRequestDetails(method, headers, params, body)

	context.JSON(http.StatusOK, details)

}
