package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
JSONHandler

HttpBin /json davranışı:

1) Endpoint:
   GET /json

2) Response:
   - Sabit bir JSON döner
   - İçeriği önemli ama genelde:
     {
       "slideshow": {
         "author": "...",
         "date": "...",
         "slides": [
           {
             "title": "...",
             "type": "..."
           }
         ]
       }
     }

3) Status code:
   - 200 OK

4) Header:
   - Content-Type: application/json

5) Request body / query / header:
   - ÖNEMSİZ
   - Ne gelirse gelsin aynı JSON döner

6) Streaming yok
7) Compression yok
*/

type Response struct {
	AuthorName    string `json:"author_name"`
	AuthorSurname string `json:"author_surname"`
	AuthorGithub  string `json:"author_github"`
	AuthorAge     int    `json:"author_age"`
}

func JSONHandler(c *gin.Context) {
	// 1) JSON objesini (map / struct) hazırla
	//    - nested map veya struct kullanabilirsin
	author := Response{
		AuthorName:    "Elif",
		AuthorSurname: "Erden",
		AuthorGithub:  "https://github.com/eliferdentr",
		AuthorAge:     26,
	}
	c.Header("Content-Type", "application/json")
	// 2) Status = 200
	c.Status(http.StatusOK)
	// 3) c.JSON(200, obj) ile response yaz
	c.JSON(http.StatusOK, author)
}
