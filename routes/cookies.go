package routes

import (
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/controllers"
)

func RegisterCookiesRoutes(router *gin.Engine) {
	router.GET("/cookies", controllers.GetCookies)
	router.GET("/cookies/set", controllers.SetCookies)
	router.GET("/cookies/delete", controllers.DeleteCookies)
}