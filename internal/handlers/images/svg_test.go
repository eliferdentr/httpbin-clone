package images

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSVGHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/image/svg", SVGHandler)

	req := httptest.NewRequest("GET", "/image/svg", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	if rec.Header().Get("Content-Type") != "image/svg+xml" {
		t.Fatalf("expected image/svg+xml")
	}

	if rec.Body.Len() == 0 {
		t.Fatal("empty svg body")
	}
}
