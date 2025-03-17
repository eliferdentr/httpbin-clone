package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils"
)

//Sending n JSON messages as a stream over a single HTTP connection,
//  with a short delay in between (e.g. 500ms), up to the specified n value.

func StreamHandler (context *gin.Context) {
	n, err := utils.GetPositiveIntParam(context, "n")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	count := 0

	

}