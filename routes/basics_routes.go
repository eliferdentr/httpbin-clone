package routes

import (
	"github.com/gin-gonic/gin"
	controller "httbinclone-eliferden.com/controllers"
)

func RegisterBasicRoutes(router *gin.Engine) {
	router.GET("/", controller.RootHandler)
	router.GET("/uuid", controller.GetUUID)
}
