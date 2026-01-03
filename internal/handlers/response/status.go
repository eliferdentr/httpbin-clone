package response

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
1. URL parametresi = döndürülecek HTTP status code

Örn:
/status/418 → “I’m a teapot”
/status/404
/status/204

2. Body göndermez
Yani:
Status 204 → body tamamen boş
Status 200 → body yine boş
Her durum için sadece raw status code döner.

3. Geçersiz kod → yine o kod ile döner

Yani numeric valid range olduğu sürece problem yok.

4. Content-Type yok

c.Status(code) kullanarak hızlıca çözebilirsin.
*/

// StatusHandler returns the given HTTP status code.
//
// @Summary Return a status code
// @Description Returns the given status code with empty body (except teapot).
// @Tags response
// @Param code path int true "HTTP status code" minimum(100) maximum(599)
// @Success 204 {string} string "No Content"
// @Success 200 {string} string "OK"
// @Success 418 {string} string "I'm a teapot"
// @Failure 400 {object} map[string]string
// @Router /status/{code} [get]
func StatusHandler(c *gin.Context) {
	codeStr := c.Param("code")
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if code < 100 || code > 599 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status code must be between 100 and 599"})
		return
	}
	if code == http.StatusTeapot {
		c.String(http.StatusTeapot, "I'm a teapot")
		return
	}
	c.Status(code)
}
