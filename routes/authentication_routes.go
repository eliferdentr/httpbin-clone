package routes

import (
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/controllers"
)

func RegisterAuthenticationRoutes(router *gin.Engine) {
	router.GET("/basic-auth/:user/:password", controllers.VerifyBasicAuth)
	router.GET("/bearer", controllers.VerifyBearerAuth)
	router.GET("/digest-auth/:qop/:user/:passwd",controllers.VerifyDigestAuthTypes)
	router.GET("/digest-auth/:qop/:user/:passwd/:algorithm",controllers.VerifyDigestAuthTypes)
	router.GET("/digest-auth/:qop/:user/:passwd/:algorithm/:stale_after",controllers.VerifyDigestAuthTypes)
}