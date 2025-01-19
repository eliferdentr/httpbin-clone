package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func PostUpload(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to upload the file. Please upload a valid file",
		})
		return
	}

	fileExtension := filepath.Ext(file.Filename)
	context.JSON(http.StatusOK, gin.H{
		"message" : "File extention extracted successfully",
		"file_name" : file.Filename,
		"file_Extention" : fileExtension,
	})
}