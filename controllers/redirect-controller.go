package controllers

import (
	"fmt"
	"net/http"
	"strconv"
    constants "httbinclone-eliferden.com/utils/constants"
	"github.com/gin-gonic/gin"
)

//redirects for n times to the same endpoint
func RedirectHandler(context *gin.Context) {
    n, err := getPositiveIntParam(context, "n")
	if err != nil {
		return // Hata varsa yanıt döndürüldü, devam etmiyoruz
	}

    if n > 1 {
        redirectUrl := fmt.Sprintf("/redirect/%d", n-1)
        //redirect to the same endpoint
        context.Redirect(http.StatusFound, redirectUrl)
        return
    }
    context.JSON(http.StatusBadRequest, gin.H{
        "message" : "Final destination reached after redirects.",
    })

}

//redirects with absolute path for n times
func AbsoluteRedirectHandler(context *gin.Context) {
    n, err := getPositiveIntParam(context, "n")
	if err != nil {
		return // Hata varsa yanıt döndürüldü, devam etmiyoruz
	}

    if n > 1 {
        redirectUrl := fmt.Sprintf("%s/absolute-redirect/%d",constants.BasePath,  n-1)
        //redirect to the same endpoint
        context.Redirect(http.StatusFound, redirectUrl)
        return
    }
    context.JSON(http.StatusBadRequest, gin.H{
        "message" : "Final destination reached after redirects.",
    })
}

//redirects with relative path for n times
func RelativeRedirectHandler(context *gin.Context) {
	n, err := getPositiveIntParam(context, "n")
	if err != nil {
		return // Hata varsa yanıt döndürüldü, devam etmiyoruz
	}

	// Göreceli URL oluştur
	if n > 1 {
		redirectURL := fmt.Sprintf("/relative-redirect/%d", n-1)
		context.Redirect(http.StatusFound, redirectURL)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Final destination reached after relative redirects.",
	})
}

func getPositiveIntParam(context *gin.Context, paramName string) (int, error) {
	paramStr := context.Param(paramName) // Parametreyi al
	param, err := strconv.Atoi(paramStr) // String to int dönüşümü
	if err != nil || param <= 0 {        // Geçersizlik kontrolü
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter. Please provide a positive integer.",
		})
		return 0, err
	}
	return param, nil
}