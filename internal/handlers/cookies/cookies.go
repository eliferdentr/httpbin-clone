package cookies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /cookies
// Tüm cookie’leri JSON olarak döndürür
func CookiesHandler(c *gin.Context) {
	// 1) request'ten cookie map oluştur
	cookies := createCookieMap(c.Request.Cookies())
	// 2) JSON olarak {"cookies": {...}} döndür
	
}

// GET /cookies/set/:name/:value
// Cookie set eder → sonra redirect (/cookies)
func SetCookieHandler(c *gin.Context) {
	// 1) name, value parametrelerini al
	// 2) cookie oluştur → HttpOnly=false, Path="/"
	// 3) c.SetCookie(...)
	// 4) redirect → 302 → /cookies
}

// GET /cookies/delete
// Bir veya daha fazla cookie'yi siler → redirect (/cookies)
// /cookies/delete?name=a&name=b
func DeleteCookieHandler(c *gin.Context) {
	// 1) c.QueryArray("name") ile cookie isimlerini al
	// 2) her biri için:
	//      maxAge= -1 vererek cookie sil
	// 3) redirect → 302 → /cookies
}

func createCookieMap([]*http.Cookie) map[string]*http.Cookie {
	cookiesMap := make(map[string]*http.Cookie)
	for _, cookie := range cookiesMap {
		cookiesMap[cookie.Name] = cookie
	}
	return cookiesMap
}
