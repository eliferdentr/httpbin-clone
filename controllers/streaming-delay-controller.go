package controllers

import (
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils"
)

//Sending n JSON messages as a stream over a single HTTP connection,
//  with a short delay in between (e.g. 500ms), up to the specified n value.

func GetStream (context *gin.Context) {
	n, err := utils.GetPositiveIntParam(context, "n")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	count := 0
	//context.Stream, calls the function written inside it repeatedly
	//the function inside sends chunk data in every call
	//return value of the function needs to be bool:
		//if it returns true then Gin continues streaming (tries to send more chunk)
		//if it returns false then streaming stops

		//send a message using server-sent events
		context.Stream(func(w io.Writer) bool {
			if count < n {
				// (Server-Sent Events)
				context.SSEvent("message", gin.H{
					"id":      count,
					"message": fmt.Sprintf("Streamed message %d", count),
				})
				count++
				time.Sleep(500 * time.Millisecond) 
				return true
			}
			return false 
		})

}

func GetDelay (context *gin.Context) {
	seconds, err := utils.GetPositiveIntParam(context, "n")
	if  err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	time.Sleep(time.Duration(seconds) * time.Second)

	context.JSON(http.StatusOK, gin.H{
		"message" : fmt.Sprintf("Response delayed by %d seconds", seconds),
		"delay_seconds" : seconds,
	})
}


func GetDrip (context *gin.Context) {
	totalChunks, err := utils.GetPositiveIntParam(context, "n")
	if  err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	interval := 1 * time.Second
	chunkData := "0123456789"

	count := 0
	context.Stream(func(w io.Writer) bool {
		if count < totalChunks {
			// Yaz, ardından buffer'ı flush et
			context.Writer.Write([]byte(chunkData))
			context.Writer.Flush()
			count++
			time.Sleep(interval)
			return true
		}
		return false // Tüm chunk'lar gönderildikten sonra stream'i kapat
	})

}