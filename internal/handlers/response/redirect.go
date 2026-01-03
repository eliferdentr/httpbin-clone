package response

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
Redirect dünyasında 3 farklı endpoint var ve hepsi aslında aynı mantığın varyasyonu:

HttpBin Redirect Ailesi
1- /redirect/:n
Toplam n kez redirect yapar.
Redirect path’i /get ile biter (HttpBin böyle yapıyor).
Redirect’ler absolute URL şeklindedir.
Örnek:
/redirect/3
→ 302 → /redirect/2
→ 302 → /redirect/1
→ 302 → /get
→ 200 → JSON

2- /relative-redirect/:n
Aynı mantık.
Ama redirect header’daki Location değeri relative path olur:
Örneğin:
Location: /relative-redirect/2

3- /absolute-redirect/:n
Aynı mantık.
Ama bütün redirect'lerde Location header’ı tam URL içerir:
Location: http://example.com/absolute-redirect/4
HttpBin bunu request.Host üzerinden oluşturur.
*/

// RedirectHandler godoc
//
// @Summary      Redirect N times
// @Description Performs N redirects ending at /get
// @Tags         response
// @Param        n path int true "Redirect count"
// @Success      302
// @Failure      400 {object} map[string]string
// @Router       /redirect/{n} [get]
func RedirectHandler(c *gin.Context) {
	// n alınır
	nStr := c.Param("n")
	if nStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "n is required"})
		return
	}
	// n <= 0 ise /get’e gönderilir
	n, err := strconv.Atoi(nStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "n is invalid"})
		return
	}
	if n <= 0 {
		c.Redirect(http.StatusFound, "/get")
	} else {
		// aksi halde:
		//   Location: /redirect/(n-1)
		//   c.Redirect(302, url)
		newNStr := strconv.Itoa(n - 1)
		newRedirect := "/redirect/" + newNStr
		c.Redirect(http.StatusFound, newRedirect)
	}
}

// RelativeRedirectHandler godoc
//
// @Summary      Relative redirect
// @Description Performs relative redirects
// @Tags         response
// @Param        n path int true "Redirect count"
// @Success      302
// @Failure      400 {object} map[string]string
// @Router       /relative-redirect/{n} [get]
func RelativeRedirectHandler(c *gin.Context) {
	nStr := c.Param("n")
	if nStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "n is required"})
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "n is invalid"})
		return
	}

	if n <= 0 {
		c.Redirect(http.StatusFound, "/get")
		return
	}

	newN := strconv.Itoa(n - 1)
	c.Redirect(http.StatusFound, "/relative-redirect/"+newN)
}

// AbsoluteRedirectHandler godoc
//
// @Summary      Absolute redirect
// @Description Performs absolute redirects
// @Tags         response
// @Param        n path int true "Redirect count"
// @Success      302
// @Failure      400 {object} map[string]string
// @Router       /absolute-redirect/{n} [get]
func AbsoluteRedirectHandler(c *gin.Context) {
	nStr := c.Param("n")
	if nStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "n is required"})
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "n is invalid"})
		return
	}

	requestHost := c.Request.Host

	if n <= 0 {
		c.Redirect(http.StatusFound, "http://"+requestHost+"/get")
		return
	}

	newN := strconv.Itoa(n - 1)
	newURL := "http://" + requestHost + "/absolute-redirect/" + newN

	c.Redirect(http.StatusFound, newURL)
}
