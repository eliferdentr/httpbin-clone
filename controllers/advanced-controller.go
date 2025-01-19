package controllers

import "github.com/gin-gonic/gin"

func GetAnything(context *gin.Context) {
	//return details of the request
	headers := context.Request.Header
	method := context.Request.Method
	url := context.Request.URL.String()
	params :=


}

func GetDelay (context *gin.Context) {
	
}