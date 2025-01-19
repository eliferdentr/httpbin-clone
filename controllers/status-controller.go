package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStatus(context *gin.Context) {
	statusCode := context.Param("code")
	if statusCode == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error" : "The param 'code' has not been provided",
		})
		return
	}
	code, err := strconv.Atoi(statusCode)
	if err != nil || code < 100 || code > 599 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid status code. Please provide a valid HTTP status code (100-599).",
		})
		return
	}
	context.JSON(code, gin.H{
		"message": "This is your custom status code response.",
		"status":  code,
	})
}