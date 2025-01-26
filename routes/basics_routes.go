package routes

import (
	"github.com/gin-gonic/gin"
	controller "httbinclone-eliferden.com/controllers"
)

func RegisterBasicRoutes(router *gin.Engine) {
	router.GET("/", controller.RootHandler)
	router.GET("/ip", controller.GetIP)
	router.GET("/uuid", controller.GetUUID)
	router.GET("/user-agent", controller.GetUserAgent)
	router.GET("/headers", controller.GetHeaders)
}
