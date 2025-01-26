package routes

import (
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/controllers"
)

func RegisterStatusRoutes(router *gin.Engine) {
	router.GET("/status/:code", controllers.GetStatus)
}