package routes
import (
	"github.com/gin-gonic/gin"
)

func RegisterBasicRoutes (router *gin.Engine) {
router.GET("/get")
router.POST("/POST")
}