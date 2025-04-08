package controllers

import (
	"compress/gzip"
	"compress/zlib"
	"crypto/rand"
	"encoding/json"
	"io"
	"net/http"

	"github.com/andybalholm/brotli"
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/assets"
	services "httbinclone-eliferden.com/services/implementation"
	"httbinclone-eliferden.com/utils"
	"httbinclone-eliferden.com/utils/constants"
)

//generates and returns random data of incoming bytes
func GetRange(context *gin.Context){
	n, err := utils.GetPositiveIntParam(context, "n")
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	data := make([]byte, n)
	_, err = rand.Read(data)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate random data.",
		})
		return
	}

	// Veriyi application/octet-stream tipinde döndür
	context.Data(http.StatusOK, "application/octet-stream", data)
}

//return an html page
func GetHTML( context *gin.Context){
	htmlContent := constants.HTMLCONTENT
	if htmlContent == "" {
		context.JSON(http.StatusInternalServerError,
			gin.H{"error": `There was an error while handling the request.
			HTML Template could not be found on the server.`})
			return
	}
	context.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
}

//return scanning rules of the site
func GetRobots ( context *gin.Context) {
	robotsContent := constants.ROBOTSCONTENT
	if robotsContent == "" {
		context.JSON(http.StatusInternalServerError,
			gin.H{"error": `There was an error while handling the request.
			robots.txt could not be found on the server.`})
			return
	}
	context.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(robotsContent))
}

//stops the incoming request and denies it
func GetDeny (context *gin.Context) {
context.AbortWithStatusJSON(http.StatusForbidden, gin.H{
	"error" : "Access Denied",
})
}

//get the embedded image
func GetImage (context *gin.Context) {
	context.Data(http.StatusOK, "image/jpeg", assets.SampleImage)
}


func PostForms(context *gin.Context) { 
	formContent := constants.FORMCONTENT;
	if formContent == "" {
		context.JSON(http.StatusInternalServerError,
			gin.H{"error": `There was an error while handling the request.
			form could not be found on the server.`})
			return
	}
	context.Data(http.StatusOK, "text/html; charset=utf-8", []byte(formContent))
	
}

//used to send random data (for example, secure random data) to the client in chunks, totaling n bytes.
func GetStreamBytes (context *gin.Context) {
	n, err := utils.GetPositiveIntParam(context, "n")
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
			return
	}
	context.Header("Content-Type", "application/octet-stream")
	
	writtenBytes := 0

	//we genearate 1024 byte chunk every time and we keep this chunk in buffer
	//at first, then we send it to the client.
	 
	context.Stream(func(w io.Writer) bool {
		if writtenBytes >= n {
			return false
		}
		//send 1024 byte chunks every time
		chunkSize := 1024
		remain := n - writtenBytes

		if remain < chunkSize {
			chunkSize = remain
		}

		buffer := make([]byte, chunkSize)


		//fill in with random bytes
		if _, err := rand.Read(buffer); err != nil {
			return false
		}

		//write bytes
		if _, err := w.Write(buffer); err != nil {
			return false
		}
		writtenBytes += chunkSize
		return true
	})

}

func GetGZip (context *gin.Context) {
	context.Header("Content-Encoding", "gzip")
	context.Header("Content-Type", "application/json")

	//create a new gzzip writer
	gw := gzip.NewWriter(context.Writer)
	defer gw.Close()

	//the response we want to send
	response := map[string]interface{} {
		"gzipped":true,
		"method": context.Request.Method,
	}

	//create the json data
	jsonData, err := json.Marshal(response)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate json: " + err.Error(),
		})
		return
	}
	if _, err := gw.Write(jsonData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to write the compressed data: " + err.Error(),
		})
		return
	}

}

//another way to compress the data
func GetDeflate (context * gin.Context) {
	context.Header("Content-Encoding", "deflate")
	context.Header("Content-Type", "application/json")
	// zlib kullanarak deflate writer oluşturuyoruz (zlib, deflate algoritmasını kullanıyor)
	zw := zlib.NewWriter(context.Writer)
	defer zw.Close()

	// Göndermek istediğimiz JSON verisi
	response := map[string]interface{}{
		"deflated": true,
		"method":   context.Request.Method,
	}

	// JSON verisini marshal ediyoruz
	jsonData, err := json.Marshal(response)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate json: " + err.Error(),
		})
		return
	}
	if _, err := zw.Write(jsonData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to write the compressed data: " + err.Error(),
		})
		return
	}
}

func GetBrotli (context *gin.Context) {
	context.Header("Content-Encoding", "br")
	context.Header("Content-Type", "application/json")

	// c.Writer'ı sardığımız Brotli writer'ı oluşturuyoruz
	bw := brotli.NewWriter(context.Writer)
	defer bw.Close()

	// Döndüreceğimiz JSON verisi
	response := map[string]interface{}{
		"brotli": true,
		"method": context.Request.Method,
	}

	// JSON verisini marshal ediyoruz
	jsonData, err := json.Marshal(response)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate json: " + err.Error(),
		})
		return
	}
	if _, err := bw.Write(jsonData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to write the compressed data: " + err.Error(),
		})
		return
	}
}

func GetWebSocket(context *gin.Context) {
	
}

func GetAnything(context *gin.Context) {

	headers := context.Request.Header
	method := context.Request.Method
	params := context.Request.URL.Query()

	var body map[string]any
	if err := context.ShouldBindJSON(&body); err != nil {
		body = nil
	}

	service := services.NewRequestProcessorServiceImpl()
	details := service.GetRequestDetails(method, headers, params, body)

	context.JSON(http.StatusOK, details)

}
