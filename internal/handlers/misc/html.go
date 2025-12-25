package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
HTMLHandler (genel mantık)

HttpBin /html davranışı:

 1. Endpoint:
    GET /html

2) Response:
  - Content-Type: text/html; charset=utf-8
  - Body: DÜZ HTML STRING

3) JSON YOK
4) Template engine YOK
5) Status: 200 OK
6) Her zaman aynı HTML döner
*/
func HTMLHandler(c *gin.Context) {

	// 1) Content-Type ayarla
	c.Header("Content-Type", "text/html; charset=utf-8")

	// 2) Status code
	c.Status(http.StatusOK)

	// 3) HTML body yaz
	//    (raw string literal kullanmak en temiz yol)
	c.String(http.StatusOK, `
<!DOCTYPE html>
<html>
<head>
	<title>HttpBin Clone</title>
</head>
<body>
	<h1>Hello World</h1>
	<p>This is a simple HTML response.</p>
</body>
</html>
`)
}
