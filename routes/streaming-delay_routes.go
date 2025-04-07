package routes

import (
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/controllers"
)

func RegisterStreamingDelayRoutes(router *gin.Engine) {
	router.GET("/stream/:n", controllers.GetStream)
	router.GET("/delay/:n", controllers.GetDelay)
	router.GET("/drip/:n", controllers.GetDrip)

}