package cookies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCookiesHandler(c *gin.Context) {
	cookies := c.Request.Cookies()
	cookiesMap := make (map[string]string)

	for _, cookie := range cookies{
		cookiesMap[cookie.Name] = cookie.Value
	}

	c.JSON(http.StatusOK, gin.H{
		"cookies": cookiesMap,
	})
}