package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//returns the list of all the endpoints
func RootHandler(context *gin.Context) {
	routes := gin.Default().Routes()
	if len(routes) < 1 {
		context.JSON(http.StatusOK, gin.H{
			"message" : "There were not any registered routes found.",
		})
		return
	}
	filteredRoutes := make([]string, len(routes))
	for _, route := range routes {
		filteredRoutes = append(filteredRoutes, route.Path)
	}
	context.JSON(http.StatusOK, gin.H{
		"message" : "There were not any registered routes found.",
	})
}

//returns a random uuid
func GetUUID(context *gin.Context) {
	uuidString := uuid.NewString()
	context.JSON(http.StatusOK, gin.H {
		"generated_uuid" : uuidString,
	})
}


