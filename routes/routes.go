package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	//advanced
	RegisterAdvancedRoutes(router)
	
	//headers
	RegisterHeadersRoutes(router)

	//basics
	RegisterBasicRoutes(router)

	//status
	RegisterStatus(router)

	//upload
	RegisterUploadRoutes(router)

}