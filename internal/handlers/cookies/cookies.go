package cookies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /cookies
// Tüm cookie’leri JSON olarak döndürür

// @Summary      Get cookies
// @Description  Returns all cookies sent by the client
// @Tags         cookies
// @Produce      application/json
// @Success      200 {object} map[string]map[string]string
// @Router       /cookies [get]
// @Failure      500 {object} map[string]string
func CookiesHandler(c *gin.Context) {
	// 1) request'ten cookie map oluştur
	cookies := c.Request.Cookies()
	cookiesMap := make(map[string]string)
	for _, cookie := range cookies {
		cookiesMap[cookie.Name] = cookie.Value
	}
	// 2) JSON olarak {"cookies": {...}} döndür
	c.JSON(http.StatusOK, gin.H{
		"cookies": cookiesMap,
	})

}

// GET /cookies/set/:name/:value
// Cookie set eder → sonra redirect (/cookies)

// @Summary      Set a cookie
// @Description  Sets a cookie and redirects to /cookies
// @Tags         cookies
// @Param        name   path string true "Cookie name"
// @Param        value  path string true "Cookie value"
// @Success      302
// @Header       302 {string} Location "/cookies"
// @Router       /cookies/set/{name}/{value} [get]
// @Failure      500 {object} map[string]string
func SetCookieHandler(c *gin.Context) {
	// 1) name, value parametrelerini al
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is empty"})
	}
	value := c.Param("value")
	if value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "value is empty"})
	}
	// 2) cookie oluştur → HttpOnly=false, Path="/"
	// 3) c.SetCookie(...)
	c.SetCookie(name, value, 0, "/", "", false, false)
	// 4) redirect → 302 → /cookies
	c.Redirect(http.StatusFound, "/cookies")
}

// @Summary      Delete cookies
// @Description  Deletes one or more cookies and redirects to /cookies
// @Tags         cookies
// @Param        name  query []string true "Cookie name(s) to delete"
// @Success      302
// @Header       302 {string} Location "/cookies"
// @Router       /cookies/delete [get]

// GET /cookies/delete
// Bir veya daha fazla cookie'yi siler → redirect (/cookies)
// /cookies/delete?name=a&name=b
func DeleteCookieHandler(c *gin.Context) {
	// 1) c.QueryArray("name") ile cookie isimlerini al
	cookieNames := c.QueryArray("name")
	// 2) her biri için:
	//      maxAge= -1 vererek cookie sil
	for _, n := range cookieNames {
		if n != "" {
			// MaxAge = -1 → delete
			c.SetCookie(n, "", -1, "/", "", false, false)
		}
	}

	c.Redirect(http.StatusFound, "/cookies")
}
