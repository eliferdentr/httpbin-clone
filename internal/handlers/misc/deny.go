package misc

import (
	// net/http import edilecek
	// gin import edilecek

	"net/http"

	"github.com/gin-gonic/gin"
)

/*
DenyHandler

HttpBin /deny davranışı:

1) Her zaman erişimi reddeder
2) HTTP Status Code: 403 Forbidden
3) Body:
   {
     "message": "YOU SHALL NOT PASS"
   }
   (string birebir önemli olabilir, teste göre ayarlanır)

4) Method, header, query fark etmez
   → GET, POST, ne gelirse gelsin aynı response

5) Ekstra header yok
6) Redirect yok
*/

// DenyHandler godoc
//
// @Summary      Always deny access
// @Description Always returns 403 Forbidden
// @Tags         misc
// @Produce      application/json
// @Success      403 {object} map[string]string
// @Router       /deny [get]
func DenyHandler(c *gin.Context) {
	// 1) Status code'u 403 olarak ayarla
	// 2) JSON body yaz
	//    message alanı ile
	c.JSON(http.StatusForbidden, gin.H{
		"message": "YOU SHALL NOT PASS",
	})
	return
}
