package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestBytesHandler_ValidSize(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/bytes/:n", BytesHandler)
	// 3) request: /bytes/10
	req := httptest.NewRequest("GET", "/bytes/10", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) status = 200
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}
	// 8) content-type = application/octet-stream mı?
	contentTypeHeader := recorder.Header().Get("Content-Type")
	if contentTypeHeader == "" {
		t.Fatalf("Content-Type cannot be empty")
	}
	if contentTypeHeader != "application/octet-stream" {
		t.Fatalf("Expected Content-Type is application/octet-stream. Received: %q", contentTypeHeader)
	}
	// 9) body length = 10 byte mı?
	if recorder.Body.Len() != 10 {
		t.Fatalf("Expected length is 10. Received: %d", recorder.Body.Len())
	}

}

func TestBytesHandler_InvalidNumber(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/bytes/:n", BytesHandler)
	// 3) request: /bytes/abc" gönder
	req := httptest.NewRequest("GET", "/bytes/abc", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 2) status 400 bekle
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected code is 400. Received: %d", recorder.Code)
	}
}

func TestBytesHandler_NegativeNumber(t *testing.T) {
		// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/bytes/:n", BytesHandler)
	// 3) request: "/bytes/-5" gönder
	req := httptest.NewRequest("GET", "/bytes/-1", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 2) status 400 bekle
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected code is 400. Received: %d", recorder.Code)
	}
}

func TestBytesHandler_ZeroBytes(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/bytes/:n", BytesHandler)
	// 3) request: /bytes/0
	req := httptest.NewRequest("GET", "/bytes/0", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) status = 200
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}
	// 9) body length = 0 mı?
	if recorder.Body.Len() != 0 {
		t.Fatalf("Expected length is 0. Received: %d", recorder.Body.Len())
	}
}
