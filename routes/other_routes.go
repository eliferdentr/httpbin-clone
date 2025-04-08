package routes

import (
	"github.com/gin-gonic/gin"
	controller "httbinclone-eliferden.com/controllers"
)

func RegisterOtherRoutes(router *gin.Engine) {
	router.GET("/range/:n", controller.GetRange)
	router.GET("/html", controller.GetHTML)
	router.GET("/robots.txt", controller.GetRobots)
	router.GET("/deny", controller.GetDeny)
	router.GET("/image", controller.GetImage)
	router.GET("/forms/post", controller.PostForms)
	router.GET("/stream-bytes/:n", controller.GetStreamBytes)
	router.GET("/gzip", controller.GetGZip)
	router.GET("/deflate", controller.GetDeflate)
	router.GET("/brotli", controller.GetBrotli)
	router.GET("/anything", controller.GetAnything)
	router.GET("/range/:n", controller.GetRange)
	router.GET("/websocket", controller.GetWebSocket)

}
