package routes
import (
	"github.com/gin-gonic/gin"
	controller "httbinclone-eliferden.com/controllers"
)

func RegisterResponseInspectionRoutes(router *gin.Engine) {
	router.GET("/cache", controller.CacheHandler)
	router.GET("/cache/:value", controller.CacheWithValueHandler)
	router.GET("/etag/:etag", controller.ETagHandler)
	router.GET("/response-headers", controller.ResponseHeadersForGetHandler)
	router.POST("/response-headers", controller.ResponseHeadersForPostHandler)
}