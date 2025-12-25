package misc

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRobotsHandler(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)

	// 2) router
	r := gin.Default()
	r.GET("/robots.txt", RobotsHandler)

	// 3) request
	req := httptest.NewRequest("GET", "/robots.txt", nil)
	recorder := httptest.NewRecorder()

	// 4) serve
	r.ServeHTTP(recorder, req)

	// 5) status code = 200
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}

	// 6) Content-Type = text/plain
	if recorder.Header().Get("Content-Type") != "text/plain" {
		t.Fatalf(
			"expected Content-Type text/plain, got %s",
			recorder.Header().Get("Content-Type"),
		)
	}

	// 7) body birebir kontrol
	expectedBody := "User-agent: *\nDisallow: /deny"
	if recorder.Body.String() != expectedBody {
		t.Fatalf(
			"expected body:\n%q\ngot:\n%q",
			expectedBody,
			recorder.Body.String(),
		)
	}
}
