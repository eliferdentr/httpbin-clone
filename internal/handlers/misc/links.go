package misc

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
LinksHandler

HttpBin /links/:n davranışı:

1) Endpoint:
   GET /links/:n

2) n adet link üretir:
   http://example.com/links/0
   http://example.com/links/1
   ...

3) Response:
{
  "links": [
    "http://example.com/links/0",
    ...
  ]
}

4) Status: 200
5) Content-Type: application/json
6) n negatif veya sayı değilse: 400
*/

// LinksHandler godoc
//
// @Summary      Generate links
// @Description Generates N HTML links
// @Tags         misc
// @Produce      text/html
// @Param        n path int true "Number of links"
// @Success      200 {string} string "HTML links"
// @Failure      400 {object} map[string]string
// @Router       /links/{n} [get]
func LinksHandler(c *gin.Context) {
	nStr := c.Param("n")

	n, err := strconv.Atoi(nStr)
	if err != nil || n < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid n",
		})
		return
	}

	links := make([]string, 0, n)
	host := c.Request.Host
	if host == "" {
		host = "example.com"
	}

	for i := 0; i < n; i++ {
		link := "http://" + host + "/links/" + strconv.Itoa(i)
		links = append(links, link)
	}

	c.JSON(http.StatusOK, gin.H{
		"links": links,
	})
}
