package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils"
	constants "httbinclone-eliferden.com/utils/constants"
)

//redirects for n times to the same endpoint
func RedirectHandler(context *gin.Context) {
    n, err := utils.GetPositiveIntParam(context, "n")
	if err != nil {
		return // Hata varsa yanıt döndürüldü, devam etmiyoruz
	}

    if n > 1 {
        redirectUrl := fmt.Sprintf("/redirect/%d", n-1)
        //redirect to the same endpoint
        context.Redirect(http.StatusFound, redirectUrl)
        return
    }
    context.JSON(http.StatusOK, gin.H{
        "message" : "Final destination reached after redirects.",
    })

}

//redirects with absolute path for n times
func AbsoluteRedirectHandler(context *gin.Context) {
    n, err := utils.GetPositiveIntParam(context, "n")
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
	n, err := utils.GetPositiveIntParam(context, "n")
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
