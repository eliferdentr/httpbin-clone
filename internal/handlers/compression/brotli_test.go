package compression

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andybalholm/brotli"
	"github.com/gin-gonic/gin"
)

func TestBrotliHandler_NoAcceptEncoding(t *testing.T) {
	// gin test mode
	gin.SetMode(gin.TestMode)

	// router
	r := gin.Default()
	r.GET("/brotli", BrotliHandler)

	// Accept-Encoding OLMADAN request
	req := httptest.NewRequest("GET", "/brotli", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	// status 200 mü?
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	// Content-Encoding OLMAMALI
	if recorder.Header().Get("Content-Encoding") != "" {
		t.Fatalf("expected no Content-Encoding header")
	}

	// JSON parse edilebiliyor mu?
	var data map[string]interface{}
	if err := json.Unmarshal(recorder.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	// brotli=false mi?
	if data["brotli"] != false {
		t.Fatalf("expected brotli=false, got %v", data["brotli"])
	}
}

func TestBrotliHandler_WithBrotliEncoding(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/brotli", BrotliHandler)

	req := httptest.NewRequest("GET", "/brotli", nil)
	req.Header.Set("Accept-Encoding", "br")

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	// status
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	// Content-Encoding = br mi?
	if recorder.Header().Get("Content-Encoding") != "br" {
		t.Fatalf("expected Content-Encoding br")
	}

	// Body'yi brotli ile AÇ
	reader := brotli.NewReader(bytes.NewReader(recorder.Body.Bytes()))
	decompressed, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("failed to decompress brotli body: %v", err)
	}

	// Açılan body JSON mu?
	var data map[string]interface{}
	if err := json.Unmarshal(decompressed, &data); err != nil {
		t.Fatalf("invalid json after brotli decompress: %v", err)
	}

	// brotli=true mi?
	if data["brotli"] != true {
		t.Fatalf("expected brotli=true, got %v", data["brotli"])
	}
}
