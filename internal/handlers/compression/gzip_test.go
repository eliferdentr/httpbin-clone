package compression

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGzipHandler_WithoutGzip(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/gzip", GzipHandler)
	req := httptest.NewRequest("GET", "/gzip", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	// gzip OLMAMALI
	if recorder.Header().Get("Content-Encoding") != "" {
		t.Fatalf("did not expect Content-Encoding header")
	}

	var data map[string]interface{}
	if err := json.Unmarshal(recorder.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	if data["gzipped"] != false {
		t.Fatalf("expected gzipped=false, got %v", data["gzipped"])
	}
}

func TestGzipHandler_WithGzip(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/gzip", GzipHandler)
	req := httptest.NewRequest("GET", "/gzip", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	// gzip olcak
	if recorder.Header().Get("Content-Encoding") != "gzip" {
		t.Fatalf("expected Content-Encoding=gzip, got %s",
			recorder.Header().Get("Content-Encoding"))
	}

	// gzip body'yi AÇ
	reader, err := gzip.NewReader(recorder.Body)
	if err != nil {
		t.Fatalf("failed to create gzip reader: %v", err)
	}
	defer reader.Close()

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("failed to read gzip body: %v", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(decompressed, &data); err != nil {
		t.Fatalf("invalid json after unzip: %v", err)
	}

	if data["gzipped"] != true {
		t.Fatalf("expected gzipped=true, got %v", data["gzipped"])
	}

	if data["method"] != "GET" {
		t.Fatalf("expected method GET, got %v", data["method"])
	}
}
