package controllers

import (
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils"
)

func GetHandler(context *gin.Context) {
	//get all the params of the query
	context.JSON(200, utils.BuildResponse(context, nil))

}

func PostHandler(context *gin.Context) {
	//get the request body and return it
	var requestBody map[string]any //key of a json object is a string and the value could be anything
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(400, gin.H{
			"error": "Invalid Request Body.",
		})
		return
	}

	context.JSON(200, utils.BuildResponse(context, requestBody))
}

func PutHandler(context *gin.Context) {
	//get the request body, add a field called 'updated' and return it

	var requestBody map[string]any //key of a json object is a string and the value could be anything
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(400, gin.H{
			"error": "Invalid Request Body.",
		})
		return
	}

	requestBody["updated"] = true

	context.JSON(200, utils.BuildResponse(context, requestBody))

}

func DeleteHandler(context *gin.Context) {
	context.JSON(200, utils.BuildResponse(context, nil))

}

func PatchHandler(context *gin.Context) {
	//get the request body,  and return it

	var requestBody map[string]any //key of a json object is a string and the value could be anything
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(400, gin.H{
			"error": "Invalid Request Body.",
		})
		return
	}

	context.JSON(200, utils.BuildResponse(context, requestBody))

}
