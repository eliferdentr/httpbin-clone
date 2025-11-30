package response

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
HttpBin’de /cache/:n endpoint’i şöyle çalışır:
1. n → max-age saniyesi
Yani response şu header’ı içerir:

Cache-Control: public, max-age=n
2. Body olarak basit bir JSON döner
HttpBin:
{}

3. n sayı değilse → 400
4. n negatifse → 400

*/

func CacheHandler (c *gin.Context) {
	nStr := c.Param("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n < 0 {
		c.JSON(http.StatusBadRequest, "Invalid n parameter")
		return
	}
	maxControlValue := "public, max-age=" + nStr
	c.Header("Cache-Control", maxControlValue)
	c.JSON(http.StatusOK, gin.H{})
	
}