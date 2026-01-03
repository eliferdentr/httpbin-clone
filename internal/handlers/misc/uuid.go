package misc

import (
	"net/http"

	// net/http
	// github.com/gin-gonic/gin
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
UUIDHandler

HttpBin /uuid davranışı:

1) Endpoint:
   GET /uuid

2) Response:
{
  "uuid": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}

3) UUID:
   - RFC 4122 uyumlu
   - rastgele (v4)
   - her request'te farklı olmalı

4) Status code:
   - 200 OK

5) Header:
   - Content-Type: application/json

6) Request body / query / header:
   - ÖNEMSİZ
*/

// UUIDHandler godoc
//
// @Summary      Generate UUID
// @Description Returns a random UUID (v4)
// @Tags         misc
// @Produce      application/json
// @Success      200 {object} map[string]string
// @Router       /uuid [get]
func UUIDHandler(c *gin.Context) {
	// 1) Yeni UUID üret
	//    - uuid.New()
	//    - veya uuid.NewString()
	generatedUUID := uuid.NewString()
	// 2) JSON response hazırla
	//    {
	//      "uuid": generatedUUID
	//    }
	c.JSON(http.StatusOK, gin.H{
		"uuid": generatedUUID,
	})

	// 3) Status 200 + JSON response
}
