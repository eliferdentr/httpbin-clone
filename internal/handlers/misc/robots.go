package misc

import (
	// net/http
	// gin

	"net/http"

	"github.com/gin-gonic/gin"
)

/*
RobotsHandler

HttpBin /robots.txt davranışı:

1) Endpoint:
   GET /robots.txt

2) Response:
   - Body **PLAIN TEXT** olmalı (JSON DEĞİL)
   - robots.txt formatında

Örnek body:

User-agent: *
Disallow: /deny

3) Status Code:
   - 200 OK

4) Header:
   - Content-Type: text/plain

5) Request body / header / query:
   - ÖNEMSİZ
   - Ne gelirse gelsin aynı response döner

6) Streaming yok
7) Redirect yok
8) Cookie yok

Not:
- c.String(...) kullanımı EN TEMİZ YOL
*/

func RobotsHandler(c *gin.Context) {
	// Content-Type: text/plain
	c.Header("Content-Type", "text/plain")

	// robots.txt içeriği (HttpBin uyumlu)
	robots := "User-agent: *\nDisallow: /deny"

	// Status 200 + plain text body
	c.String(http.StatusOK, robots)
}
