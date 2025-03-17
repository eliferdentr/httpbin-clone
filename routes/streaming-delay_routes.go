package routes

import "github.com/gin-gonic/gin"

func RegisterStreamingDelayRoutes(router *gin.Engine) {
	router.GET("/stream/:n")
	router.GET("/delay/:n")
	router.GET("/drip")

}