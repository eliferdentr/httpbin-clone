package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCacheHandler_Valid(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/cache/:n", CacheHandler)
	// 3) request: /cache/60
	req := httptest.NewRequest("GET", "/cache/60", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) status = 200
	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected code : %d . Received: %d", http.StatusOK, recorder.Code)
	}
	// 7) Cache-Control header = "public, max-age=60" mı?
	header := recorder.Result().Header.Get("Cache-Control")
	if header == "" || header != "public, max-age=60" {
		t.Fatalf("Expected header value : public, max-age=60 . Received: %s", header)
	}
	// 8) body = "{}" mı?
	body := recorder.Body.String()
	if body != "{}" {
		t.Fatalf("Expected body '{}', got: %q", body)
	}

}

func TestCacheHandler_InvalidNumber(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/cache/:n", CacheHandler)
	// 3) request: /cache/abc
	req := httptest.NewRequest("GET", "/cache/abc", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) status = 400
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected code : %d . Received: %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestCacheHandler_NegativeNumber(t *testing.T) {
		// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/cache/:n", CacheHandler)
	// 3) request: /cache/-1
	req := httptest.NewRequest("GET", "/cache/-1", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) status = 400
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected code : %d . Received: %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestCacheHandler_Zero(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/cache/:n", CacheHandler)
	// 3) request: /cache/0
	req := httptest.NewRequest("GET", "/cache/0", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) status = 200
	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected code : %d . Received: %d", http.StatusOK, recorder.Code)
	}
	// 7) Cache-Control header = "public, max-age=0" mı?
	header := recorder.Result().Header.Get("Cache-Control")
	if header == "" || header != "public, max-age=0" {
		t.Fatalf("Expected header value : public, max-age=0 . Received: %s", header)
	}
	// 8) body = "{}" mı?
	body := recorder.Body.String()
	if body != "{}" {
		t.Fatalf("Expected body '{}', got: %q", body)
	}
}
