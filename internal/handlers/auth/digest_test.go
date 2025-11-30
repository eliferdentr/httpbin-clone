package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDigestAuthHandler_NoHeader(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/digest-auth/:qop/:user/:passwd", DigestAuthHandler)
	// 3) Authorization header OLMADAN request oluştur
	req := httptest.NewRequest("GET", "/digest-auth/auth/user/passwd", nil)
	// 4) recorder oluştur
	recorder := httptest.NewRecorder()
	// 5) router.ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) 401 döndüğünü doğrula
	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", recorder.Code)
	}
	// 7) WWW-Authenticate header’ında Digest bilgisi var mı kontrol et (realm, nonce, algorithm, qop)
	expected := `Digest realm="test", nonce="abcdef", algorithm="MD5", qop="auth"`
	if recorder.Header().Get("WWW-Authenticate") != expected {
		t.Fatalf("expected WWW-Authenticate header %q, got %q", expected, recorder.Header().Get("WWW-Authenticate"))
	}

}

func TestDigestAuthHandler_WithHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/digest-auth/:qop/:user/:passwd", DigestAuthHandler)

	req := httptest.NewRequest("GET", "/digest-auth/auth/user/passwd", nil)
	req.Header.Set("Authorization", "Digest something")

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	expected := `{"authenticated":true,"user":"user"}`
	if recorder.Body.String() != expected {
		t.Fatalf("expected body %q, got %q", expected, recorder.Body.String())
	}
}

func TestDigestAuthHandler_InvalidHeaderFormat(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/digest-auth/:qop/:user/:passwd", DigestAuthHandler)
	// 3) Authorization header'ı "Digesttt" gibi hatalı formatta gönder
	req := httptest.NewRequest("GET", "/digest-auth/auth/user/passwd", nil)
	req.Header.Set("Authorization", "Digest-broken something")
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) router.ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) Şu anki handler davranışına göre → header varsa direkt 200 dönüyor
	//    Bu yüzden expected=200 (çünkü gerçek digest doğrulaması yok)
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}
	// 7) authenticated=true kontrolü
	expected := `{"authenticated":true,"user":"user"}`
	if recorder.Body.String() != expected {
		t.Fatalf("expected body %q, got %q", expected, recorder.Body.String())
	}
}

func TestDigestAuthHandler_UserParamReturned(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/digest-auth/:qop/:user/:passwd", DigestAuthHandler)
	// 3) Authorization header'ı "Digesttt" gibi hatalı formatta gönder
	req := httptest.NewRequest("GET", "/digest-auth/auth/user1/passwd", nil)
	req.Header.Set("Authorization", "Digest something")
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) router.ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) Şu anki handler davranışına göre → header varsa direkt 200 dönüyor
	//    Bu yüzden expected=200 (çünkü gerçek digest doğrulaması yok)
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}
	// 6) Dönen JSON’da "user" parametresinin c.Param("user") ile eşleştiğini doğrula
	expected := `{"authenticated":true,"user":"user1"}`
	if recorder.Body.String() != expected {
		t.Fatalf("expected body %q, got %q", expected, recorder.Body.String())
	}
	
}
