package images

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPNGHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/image/png", PNGHandler)

	req := httptest.NewRequest("GET", "/image/png", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	if rec.Header().Get("Content-Type") != "image/png" {
		t.Fatalf("expected image/png")
	}

	if rec.Body.Len() == 0 {
		t.Fatal("empty png body")
	}
}
