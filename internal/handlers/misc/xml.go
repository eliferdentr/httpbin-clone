package misc

import (
	// net/http
	// github.com/gin-gonic/gin

	"net/http"

	"github.com/gin-gonic/gin"
)

/*
XMLHandler

HttpBin /xml davranışı:

1) Endpoint:
   GET /xml

2) Response:
   - XML formatında sabit bir doküman

Örnek (HttpBin’e benzer):

<?xml version='1.0' encoding='us-ascii'?>
<slideshow
    title="Sample Slide Show"
    date="Date of publication"
    author="Yours Truly">
  <slide type="all">
    <title>Wake up to WonderWidgets!</title>
  </slide>
  <slide type="all">
    <title>Overview</title>
    <item>Why WonderWidgets are great</item>
    <item>Who buys WonderWidgets</item>
  </slide>
</slideshow>

3) Status code:
   - 200 OK

4) Header:
   - Content-Type: application/xml

5) Body:
   - DÜZ XML string
   - JSON yok
   - Marshal şart değil (string yeterli)

6) Request body / query / header:
   - ÖNEMSİZ
*/

func XMLHandler(c *gin.Context) {
	// 1) XML string'i hazırla
	//    - backtick (`) kullanman çok rahat olur
	xml := `
<?xml version='1.0' encoding='us-ascii'?>
<slideshow
    title="Sample Slide Show"
    date="Date of publication"
    author="Yours Truly">
  <slide type="all">
    <title>Heeeeelllllooooooo!</title>
  </slide>
  <slide type="all">
    <title>Overview</title>
    <item>Why am I great</item>
    <item>Why must you think that im great nihahahaha</item>
  </slide>
</slideshow>
`
	// 2) Header:
	//    Content-Type: application/xml
	c.Header("Content-Type", "application/xml")

	// 3) Status 200
	// 4) c.String(...) ile XML body yaz
	c.String(http.StatusOK, xml)

}
