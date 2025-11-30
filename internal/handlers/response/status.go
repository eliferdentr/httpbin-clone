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

func StatusHandler(c *gin.Context) {
	codeStr := c.Param("code")
	code ,err := strconv.Atoi(codeStr)
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