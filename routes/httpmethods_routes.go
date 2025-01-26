package routes

import (
	"github.com/gin-gonic/gin"
	controller "httbinclone-eliferden.com/controllers"
)

func RegisterHTTPRoutes(router *gin.Engine) {
	router.GET("/get", controller.GetHandler)
	router.POST("/post", controller.PostHandler)
	router.PUT("/put", controller.PutHandler)
	router.DELETE("/delete", controller.DeleteHandler)
	router.PATCH("/patch", controller.PatchHandler)
}
