package cookies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCookiesHandler(c *gin.Context) {
	query := c.Request.URL.Query()
	cookies := make(map[string]string)

	for key, values := range query {
	if len(values) > 0 {
			value := values[0]
			http.SetCookie(c.Writer, &http.Cookie{
				Name:  key,
				Value: value,
				Path:  "/",
			})
			cookies[key] = value
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"cookies": cookies,
	})
}