package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	services "httbinclone-eliferden.com/services/implementation"
)

//return details of the request
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

// func GetDelay (context *gin.Context) {
// 	secondsStr := context.Param("seconds")
// 	if secondsStr == "" {
// 		context.JSON(http.StatusBadRequest, gin.H {
// 			"error" : "Missing 'seconds' parameter. Please provide a 'seconds' parameter",
// 		})
// 		return
// 	}

// 	seconds, err := strconv.Atoi(secondsStr)
// 	if err != nil || seconds < 0 {
// 		context.JSON(http.StatusBadRequest, gin.H{
// 			"error" : "MissingInvalid 'seconds' parameter. Please provide a positive integer.",
// 		})
// 		return
// 	}

// 	time.Sleep(time.Duration(seconds) * time.Second)

// 	context.JSON(http.StatusOK, gin.H{
// 		"message" : fmt.Sprintf("Response delayed by %d seconds", seconds),
// 		"delay_seconds" : seconds,
// 	})
// }