package response

import (
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

/*
1. URL’den gecikme süresi alınır

/delay/3 → 3 saniye bekle

2. Eğer sayı değilse → 400
3. Eğer saniye 0’dan küçükse → 400
4. Bekleme yapılır
*/

func DelayHandler(c *gin.Context) {
	gecikmeSuresiStr := c.Param("n")
	if gecikmeSuresiStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "n param cannot be empty"})
		return
	}
	gecikmeSuresi, err := strconv.Atoi(gecikmeSuresiStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if gecikmeSuresi < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "n parameter must be greater than or equal to 0"})
		return
	}
	time.Sleep(time.Duration(gecikmeSuresi) * time.Second)
	c.JSON(http.StatusOK, gin.H{"delay": gecikmeSuresi})

}
