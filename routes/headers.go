package routes

import "github.com/gin-gonic/gin"

func RegisterHeadersRoutes(router *gin.Engine) {
	router.GET("/headers")
	router.GET("/ip")
}