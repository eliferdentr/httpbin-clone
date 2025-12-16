package images

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestWEBPHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/image/webp", WEBPHandler)

	req := httptest.NewRequest("GET", "/image/webp", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	if rec.Header().Get("Content-Type") != "image/webp" {
		t.Fatalf("expected image/webp")
	}

	if rec.Body.Len() == 0 {
		t.Fatal("empty webp body")
	}
}
