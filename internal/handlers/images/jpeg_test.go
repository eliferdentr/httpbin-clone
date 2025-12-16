package images

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestJpegHandler_ReturnsJpegImage(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)

	// 2) router
	r := gin.Default()
	r.GET("/image/jpeg", JpegHandler)

	// 3) request
	req := httptest.NewRequest("GET", "/image/jpeg", nil)
	recorder := httptest.NewRecorder()

	// 4) serve
	r.ServeHTTP(recorder, req)

	// 5) status = 200
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	// 6) Content-Type = image/jpeg
	if recorder.Header().Get("Content-Type") != "image/jpeg" {
		t.Fatalf("expected Content-Type image/jpeg, got %s", recorder.Header().Get("Content-Type"))
	}

	// 7) body boş olmamalı
	body := recorder.Body.Bytes()
	if len(body) == 0 {
		t.Fatal("jpeg body is empty")
	}

	// 8) JPEG magic number kontrolü (FF D8)
	if body[0] != 0xFF || body[1] != 0xD8 {
		t.Fatal("response is not a valid jpeg file")
	}
}
