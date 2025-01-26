package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	//advanced
	RegisterAdvancedRoutes(router)
	//basics
	RegisterBasicRoutes(router)
	//http methods
	RegisterHTTPRoutes(router)
	//redirect
	RegisterRedirectRoutes(router)
	//status
	RegisterStatusRoutes(router)
	//cookies
	RegisterCookiesRoutes(router)

	



	//upload
	RegisterUploadRoutes(router)

}