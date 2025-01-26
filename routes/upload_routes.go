package routes

import "github.com/gin-gonic/gin"

func RegisterUploadRoutes(router *gin.Engine) {
	router.POST("/upload")
}