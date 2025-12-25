package misc

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHTMLHandler(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)

	// 2) router
	r := gin.Default()
	r.GET("/html", HTMLHandler)

	// 3) request
	req := httptest.NewRequest("GET", "/html", nil)
	recorder := httptest.NewRecorder()

	// 4) serve
	r.ServeHTTP(recorder, req)

	// 5) status code
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	// 6) Content-Type kontrolü
	ct := recorder.Header().Get("Content-Type")
	if !strings.Contains(ct, "text/html") {
		t.Fatalf("expected Content-Type text/html, got %s", ct)
	}

	// 7) Body boş olmamalı
	body := recorder.Body.String()
	if body == "" {
		t.Fatal("expected non-empty html body")
	}

	// 8) Body içinde beklenen HTML parçaları var mı?
	if !strings.Contains(body, "<html>") {
		t.Fatal("expected <html> tag in body")
	}

	if !strings.Contains(body, "<h1>") {
		t.Fatal("expected <h1> tag in body")
	}
}
