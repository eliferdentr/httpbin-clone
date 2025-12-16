package compression

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDeflateHandler_NoAcceptEncoding(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/deflate", DeflateHandler)

	req := httptest.NewRequest("GET", "/deflate", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	if rec.Header().Get("Content-Encoding") != "" {
		t.Fatalf("expected no Content-Encoding header")
	}

	var data map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	if data["deflated"] != false {
		t.Fatalf("expected deflated=false, got %v", data["deflated"])
	}
}

func TestDeflateHandler_WithDeflate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/deflate", DeflateHandler)

	req := httptest.NewRequest("GET", "/deflate", nil)
	req.Header.Set("Accept-Encoding", "deflate")

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	if rec.Header().Get("Content-Encoding") != "deflate" {
		t.Fatalf("expected Content-Encoding deflate")
	}

	// Deflate body'yi aç
	reader, err := zlib.NewReader(bytes.NewReader(rec.Body.Bytes()))
	if err != nil {
		t.Fatalf("failed to create zlib reader: %v", err)
	}
	defer reader.Close()

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("failed to read deflated body: %v", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(decompressed, &data); err != nil {
		t.Fatalf("invalid json after deflate: %v", err)
	}

	if data["deflated"] != true {
		t.Fatalf("expected deflated=true, got %v", data["deflated"])
	}
}
