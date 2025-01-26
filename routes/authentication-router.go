package routes

import (
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/controllers"
)

func RegisterAuthenticationRoutes(router *gin.Engine) {
	router.GET("/basic-auth/:user/:password", controllers.VerifyBasicAuth)
	router.GET("/hidden-basic-auth/:user/:password",controllers.VerifyHiddenBasicAuth)
	router.GET("/digest-auth/:auth/:user/:password",controllers.VerifyDigestAuth)
}