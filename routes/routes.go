package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	//basics
	RegisterBasicRoutes(router)
	//http methods
	RegisterHTTPRoutes(router)
	//status
	RegisterStatusRoutes(router)
	//redirect
	RegisterRedirectRoutes(router)
	//cookies
	RegisterCookiesRoutes(router)
	//authentication
	RegisterAuthenticationRoutes(router)
	//streaming-delay
	

	//misc


}
