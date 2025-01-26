package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils/constants"
)

//returns the cookies of the request
func GetCookies(context *gin.Context) {
	cookies, err := getAllCookies(context)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error" : err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"cookies" : cookies,
	})
}

//sets the cookies that the use has specified
func SetCookies(context *gin.Context) {
	params, err := getAllParams(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
	}

	if len(params) < 1 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error" : "No query parameters provided to set cookies.",
		})
		return
	}

	for key, value := range params {
		context.SetCookie(key, value[0], 3600, "/", constants.DOMAIN, false, true)
	}

	context.JSON(http.StatusBadRequest, gin.H{
		"message" : "Cookies are set.",
	})
}

//deletes a specified cookie
func DeleteCookies(context *gin.Context) {
	cookies, err := getAllCookies(context)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	params, err := getAllParams(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	deletedCookies := []string{}
	for key := range params {
		_, ok := cookies[key]
		if ok {
			context.SetCookie(key, "", -1, "/", "localhost", false, true)
			deletedCookies = append(deletedCookies, key)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"message":        "Cookies deleted successfully",
		"deleted_cookies": deletedCookies,
	})
}


// Helper function to get all query parameters
func getAllParams (context *gin.Context) (map[string][]string, error){
	params := context.Request.URL.Query()

	if len(params) < 1 {
		return nil, errors.New("No query parameters provided to set cookies.")
	}
	return params, nil
}
// Helper function to get all cookies
func getAllCookies(context *gin.Context) (map[string]string, error) {
	cookieSlice := context.Request.Cookies()
	if len(cookieSlice) < 1 {
		return nil, errors.New("No cookies found. Please set a cookie and try again.")
	}

	cookies := make(map[string]string)
	for _, cookie := range cookieSlice {
		cookies[cookie.Name] = cookie.Value
	}

	return cookies, nil
}


