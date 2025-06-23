package controllers

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

//returns the USER-AGENT info of the request
func GetUserAgent(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{
        "user-agent": context.Request.Header.Get("User-Agent"),
    })
}
//returns the headers
func GetHeaders(context *gin.Context) {
    // BuildResponse içindeki header oluşturma mantığını buraya taşıyalım
    headersMap := make(map[string]string)
    for k, v := range context.Request.Header {
        headersMap[k] = strings.Join(v, ", ")
    }

    context.JSON(http.StatusOK, gin.H{
        "headers": headersMap,
    })
}

//returns the client ip
func GetIP(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{
        "origin": context.ClientIP(),
    })
}
