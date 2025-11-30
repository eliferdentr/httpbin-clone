package auth

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestBasicAuthHandler_Success(t *testing.T) {
	// Gin test moduna al
	gin.SetMode(gin.TestMode)

	// Router oluştur ve handler bağla
	r := gin.Default()
	r.GET("/basic-auth/:user/:passwd", BasicAuthHandler)

	// Doğru credential için base64 header hazırla
	req := httptest.NewRequest("GET", "/basic-auth/alice/secret", nil)
	alicesecretEncoded := base64.StdEncoding.EncodeToString([]byte("alice:secret"))
	req.Header.Set("Authorization", "Basic "+alicesecretEncoded)

	// Recorder oluştur
	recorder := httptest.NewRecorder()

	//  Request’i işleme al
	r.ServeHTTP(recorder, req)

	// Beklenen durum kodu 200
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}
}

func TestBasicAuthHandler_NoHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Router oluştur ve handler bağla
	r := gin.Default()
	r.GET("/basic-auth/:user/:passwd", BasicAuthHandler)

	// Doğru credential için base64 header hazırla
	req := httptest.NewRequest("GET", "/basic-auth/alice/secret", nil)

	// Recorder oluştur
	recorder := httptest.NewRecorder()

	//  Request’i işleme al
	r.ServeHTTP(recorder, req)

	// 400 döndüğünü doğrula
	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", recorder.Code)
	}

}

func TestBasicAuthHandler_InvalidBase64(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Router oluştur ve handler bağla
	r := gin.Default()
	r.GET("/basic-auth/:user/:passwd", BasicAuthHandler)

	// TODO: Authorization: Basic !!! gibi bozuk bir base64 gönder
	req := httptest.NewRequest("GET", "/basic-auth/alice/secret", nil)
	req.Header.Set("Authorization", "Basic !!!!")

	// Recorder oluştur
	recorder := httptest.NewRecorder()

	//  Request’i işleme al
	r.ServeHTTP(recorder, req)

	// 401 döndüğünü doğrula
	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", recorder.Code)
	}

}

func TestBasicAuthHandler_WrongCredentials(t *testing.T) {
    gin.SetMode(gin.TestMode)

    r := gin.Default()
    r.GET("/basic-auth/:user/:passwd", BasicAuthHandler)

    // alice:secret istenirken alice:wrong gönder
    encoded := base64.StdEncoding.EncodeToString([]byte("alice:wrong"))
    req := httptest.NewRequest("GET", "/basic-auth/alice/secret", nil)
    req.Header.Set("Authorization", "Basic "+encoded)

    recorder := httptest.NewRecorder()
    r.ServeHTTP(recorder, req)

    if recorder.Code != http.StatusUnauthorized { // 401
        t.Fatalf("expected 401, got %d", recorder.Code)
    }
}

