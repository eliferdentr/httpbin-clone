package routes

import (
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/controllers"
)

func RegisterRedirectRoutes(router *gin.Engine) {
	router.GET("/redirect/:n", controllers.RedirectHandler)
	router.GET("/absolute-redirect/:n", controllers.AbsoluteRedirectHandler)
	router.GET("/relative-redirect/:n", controllers.RelativeRedirectHandler)
}