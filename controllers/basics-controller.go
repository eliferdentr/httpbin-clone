package controllers

import "github.com/gin-gonic/gin"

func GetGetRoute(context *gin.Context) {
	//get all the params of the query
	queryParams := context.Request.URL.Query()
	context.JSON(200, gin.H{
		"query" : queryParams,
	})

}

func PostPostRoute(context *gin.Context) {
	//get the request body and return it
	var requestBody map[string] any //key of a json object is a string and the value could be anything
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(400, gin.H{
			"error" : "Invalid Request Body.",
		})
	}

	context.JSON(200, gin.H{
		"body" : requestBody,
	})
}

func PutPutRoute(context *gin.Context) {
	//get the request body, add a field called 'updated' and return it

	var requestBody map[string] any //key of a json object is a string and the value could be anything
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(400, gin.H{
			"error" : "Invalid Request Body.",
		})
	}

	requestBody["updated"] = true

	context.JSON(200, gin.H{
		"updated_body" : requestBody,
	})
	
}

func DeleteDeleteRoute(context *gin.Context) {
	context.JSON(200, gin.H{
		"message" : "Delete request is successful",
	})
}
