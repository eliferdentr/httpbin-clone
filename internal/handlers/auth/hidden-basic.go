package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils"
)

// HiddenBasicAuthHandler godoc
//
// @Summary      Hidden Basic Authentication
// @Description  Same as basic auth but endpoint is hidden
// @Tags         auth
// @Produce      json
// @Param        user   path  string  true  "Username"
// @Param        passwd path  string  true  "Password"
// @Success      200 {object} map[string]interface{}
// @Failure      404 {object} map[string]string
// @Router       /hidden-basic-auth/{user}/{passwd} [get]

func HiddenBasicAuthHandler(c *gin.Context) {
	/*
		Aslında Basic Auth ile birebir aynı mantıkta çalışır.
		Tek fark:
		Eğer kimlik doğrulama başarısız olursa,
		sunucu 401 Unauthorized yerine 404 Not Found döner.
		Bu sayede istemci (örneğin tarayıcı) “burada basic auth var” diye algılamaz —
		yani endpoint gizlenmiş gibi olur.
		Orijinal HttpBin’de bu endpoint, “fail” durumunda WWW-Authenticate header’ını da eklemez.
	*/

	//  user, pass, header al
	// header varsa
	authorizationHeader := c.GetHeader("Authorization")
	// header yoksa 404 dön
	if authorizationHeader == "" {
		c.JSON(404, gin.H{"error": "Unauthorized"})
		return
	}
	// "Basic " prefix var mı kontrol et
	if !strings.HasPrefix(authorizationHeader, "Basic ") {
		c.JSON(404, gin.H{"error": "Invalid Authorization header format"})
		return
	}
	// decode et
	encodedPart := strings.TrimPrefix(authorizationHeader, "Basic ")
	decodedAuth, err := utils.Base64Decode(encodedPart)
	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid base64 encoding"})
		return
	}
	//  decodedAuth != expected ise yine 404 dön
	userParam := c.Param("user")
	passParam := c.Param("passwd")
	if decodedAuth != userParam+":"+passParam {
		c.JSON(404, gin.H{"error": "Unauthorized"})
		return
	}
	// 6doğruysa 200 dön {"authenticated": true, "user": userParam}

	c.JSON(200, gin.H{
		"authenticated": true,
		"user":          userParam,
	})

}
