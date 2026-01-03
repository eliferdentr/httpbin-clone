package auth

import (
	"github.com/gin-gonic/gin"
)

/*
Basic Auth gibi “user:pass” göndermez. Onun yerine tarayıcıyla sunucu bir challenge–response oyunu oynar.

1) İstemci /digest-auth/qop/user/pass endpoint’ine gelir.
Header yok → Sunucu şöyle der:
HTTP/1.1 401 Unauthorized
WWW-Authenticate: Digest
    realm="meow",
    nonce="random123",
    qop="auth"

Bu bir “challenge”.

2) Client bu bilgileri alır → user+pass+nonce ile karma hesaplar.

Sonra tekrar şu şekilde geri gönderir:

Authorization: Digest username="user",
                realm="meow",
                nonce="random123",
                uri="/digest-auth/auth/user/pass",
                response="<hash>"

3) Sunucu gelen hash’i kendisi hesaplar:

Eğer gelen response ile kendi hesapladığı hash uyuşursa:

200 OK
{ "authenticated": true, "user": "user" }


Uyuşmazsa:

401 Unauthorized
*/

// DigestAuthHandler godoc
//
// @Summary      Digest Authentication
// @Description  Digest authentication endpoint
// @Tags         auth
// @Produce      json
// @Param        qop     path string true "Quality of Protection"
// @Param        user    path string true "Username"
// @Param        passwd path string true "Password"
// @Success      200 {object} map[string]interface{}
// @Failure      401 {object} map[string]string
// @Router       /digest-auth/{qop}/{user}/{passwd} [get]
func DigestAuthHandler(c *gin.Context) {
	userParam := c.Param("user")
	authHeader := c.GetHeader("Authorization")
	// 401 + WWW-Authenticate header
	if authHeader == "" {
		c.Header("WWW-Authenticate", `Digest realm="test", nonce="abcdef", algorithm="MD5", qop="auth"`)
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Eğer header varsa direkt başarı kabul et (şimdilik)
	c.JSON(200, gin.H{
		"authenticated": true,
		"user":          userParam,
	})
}
