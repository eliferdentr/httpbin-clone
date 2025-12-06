package response

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStreamHandler_Basic(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/stream/:n", StreamHandler)

	req := httptest.NewRequest("GET", "/stream/3", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	expectedLines := []string{
		`{"id": 0}`,
		`{"id": 1}`,
		`{"id": 2}`,
	}

	body := strings.Split(strings.TrimSpace(recorder.Body.String()), "\n")

	if len(body) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(body))
	}

	for i := range expectedLines {
		if body[i] != expectedLines[i] {
			t.Fatalf("expected %q, got %q", expectedLines[i], body[i])
		}
	}
}

func TestStreamHandler_InvalidNumber(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/stream/:n", StreamHandler)

	req := httptest.NewRequest("GET", "/stream/abc", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", recorder.Code)
	}
}

func TestStreamHandler_NegativeNumber(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/stream/:n", StreamHandler)

	req := httptest.NewRequest("GET", "/stream/-3", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for negative value, got %d", recorder.Code)
	}
}
