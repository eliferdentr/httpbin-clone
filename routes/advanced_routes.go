package routes

import (
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/controllers"
)

func RegisterAdvancedRoutes(router *gin.Engine) {
	router.GET("/anything", controllers.GetAnything)
	router.GET("/delay/:seconds", controllers.GetDelay)
}