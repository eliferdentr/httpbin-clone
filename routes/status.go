package routes

import "github.com/gin-gonic/gin"

func RegisterStatus(router *gin.Engine) {
	router.GET("/status/:code")
}