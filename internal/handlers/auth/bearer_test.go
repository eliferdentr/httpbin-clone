package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBearerAuthHandler_Success(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/bearer", BearerAuthHandler)
	// 3) doğru token ile Authorization header oluştur
	req := httptest.NewRequest("GET", "/bearer", nil)
	req.Header.Set("Authorization", "Bearer abcdef")
	// 5) recorder oluştur
	recorder := httptest.NewRecorder()
	// 4) router.ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) router.ServeHTTP
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}
}

func TestBearerAuthHandler_NoHeader(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/bearer", BearerAuthHandler)
	// 3) Authorization header OLMADAN request oluştur
	req := httptest.NewRequest("GET", "/bearer", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) router.ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) 401 döndüğünü doğrula
	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", recorder.Code)
	}
}

func TestBearerAuthHandler_InvalidPrefix(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/bearer", BearerAuthHandler)
	// 3) Authorization header'ı "Token ..." gibi yanlış prefix ile oluştur
	req := httptest.NewRequest("GET", "/bearer", nil)
	req.Header.Set("Authorization", "Token abcdef")
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) router.ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) 401 döndüğünü doğrula
	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", recorder.Code)
	}
}

func TestBearerAuthHandler_EmptyToken(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/bearer", BearerAuthHandler)
	// 3) prefix doğru ama token yanlış olacak şekilde Authorization header set et
	req := httptest.NewRequest("GET", "/bearer", nil)
	req.Header.Set("Authorization", "Bearer ")
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) router.ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) 401 döndüğünü doğrula
	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", recorder.Code)
	}
}
