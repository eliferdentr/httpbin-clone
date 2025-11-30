package auth

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHiddenBasicAuthHandler_Success(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/hidden-basic-auth/:user/:passwd", HiddenBasicAuthHandler)
	// 3) doğru user/pass için base64 oluştur
	encodedUserPsswd := base64.StdEncoding.EncodeToString([]byte("user:pass"))
	// 4) request oluştur
	req := httptest.NewRequest("GET", "/hidden-basic-auth/user/pass", nil)
	req.Header.Set("Authorization", "Basic "+encodedUserPsswd)
	// 5) recorder oluştur
	recorder := httptest.NewRecorder()
	// 6) ServeHTTP yap
	r.ServeHTTP(recorder, req)
	// 7) status: 200
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}
	// 8) authenticated=true ve user=... body doğrulaması
	expected := `{"authenticated":true,"user":"user"}`
	if recorder.Body.String() != expected {
		t.Fatalf("expected body %q, got %q", expected, recorder.Body.String())
	}
}

func TestHiddenBasicAuthHandler_NoHeader(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/hidden-basic-auth/:user/:passwd", HiddenBasicAuthHandler)
	// 3) Authorization header OLMADAN request gönder
	req := httptest.NewRequest("GET", "/hidden-basic-auth/user/pass", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) status: 404 (hidden basic → her fail = 404)
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", recorder.Code)
	}
}

func TestHiddenBasicAuthHandler_InvalidPrefix(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/hidden-basic-auth/:user/:passwd", HiddenBasicAuthHandler)
	// 3) Authorization header = "Token ..." gönder (yanlış prefix)
	req := httptest.NewRequest("GET", "/hidden-basic-auth/user/pass", nil)
	req.Header.Set("Authorization", "Token abcdef")
	recorder := httptest.NewRecorder()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) status: 404 (çünkü fail → always 404)
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", recorder.Code)
	}
}

func TestHiddenBasicAuthHandler_InvalidBase64(t *testing.T) {
    // 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/hidden-basic-auth/:user/:passwd", HiddenBasicAuthHandler)
	// 3) Authorization header = "Basic !!!" gönder
	req := httptest.NewRequest("GET", "/hidden-basic-auth/user/pass", nil)
	req.Header.Set("Authorization", "Basic !!!")
	// 5) recorder oluştur
	recorder := httptest.NewRecorder()
	// 6) ServeHTTP yap
	r.ServeHTTP(recorder, req)
	// 7) status: 200
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", recorder.Code)
	}

}

func TestHiddenBasicAuthHandler_WrongCredentials(t *testing.T) {
	 // 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/hidden-basic-auth/:user/:passwd", HiddenBasicAuthHandler)
	// 3) Authorization header = doğru format + yanlış şifre
	req := httptest.NewRequest("GET", "/hidden-basic-auth/user/pass", nil)
	yanlisEncoded := base64.StdEncoding.EncodeToString([]byte("user:wrong"))
	req.Header.Set("Authorization", "Basic " + yanlisEncoded)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	// 7) status: 200
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", recorder.Code)
	}
}
