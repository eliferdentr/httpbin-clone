package routes
import (
	"github.com/gin-gonic/gin"
	controller "httbinclone-eliferden.com/controllers"
)

func RegisterRequestInspectionRoutes(router *gin.Engine) {
	router.GET("/ip", controller.GetIP)
	router.GET("/user-agent", controller.GetUserAgent)
	router.GET("/headers", controller.GetHeaders)
}