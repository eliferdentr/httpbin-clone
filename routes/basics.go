package routes
import (
	"github.com/gin-gonic/gin"
	controller "httbinclone-eliferden.com/controllers"
)

func RegisterBasicRoutes (router *gin.Engine) {
router.GET("/get", controller.GetGetRoute)
router.POST("/post", controller.PostPostRoute)
router.PUT("/put", controller.PutPutRoute)
router.DELETE("/delete", controller.DeleteDeleteRoute)
}