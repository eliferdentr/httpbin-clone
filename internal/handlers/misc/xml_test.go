package misc

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestXMLHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/xml", XMLHandler)

	req := httptest.NewRequest("GET", "/xml", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	// Status code
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	// Content-Type
	if !strings.Contains(recorder.Header().Get("Content-Type"), "application/xml") {
		t.Fatalf("expected Content-Type application/xml, got %s", recorder.Header().Get("Content-Type"))
	}

	// Body boş olmamalı
	body := recorder.Body.String()
	if body == "" {
		t.Fatal("expected non-empty XML body")
	}

	// Basit XML doğrulamaları
	if !strings.Contains(body, "<slideshow") {
		t.Fatal("expected slideshow tag in XML")
	}
	if !strings.Contains(body, "<slide") {
		t.Fatal("expected slide tag in XML")
	}
}
