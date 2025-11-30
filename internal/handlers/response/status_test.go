package response

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
)



func TestStatusHandler_ReturnsGivenCode(t *testing.T) {
    // 1) gin test mode
	gin.SetMode(gin.TestMode)
    // 2) router oluştur
	r := gin.Default()
	r.GET("/status/:code", StatusHandler)
    // 3) örneğin /status/201 isteği oluştur
	req := httptest.NewRequest("GET", "/status/201", nil)
    // 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
    // 5) 201 dönüyor mu kontrol et
	if recorder.Code !=  http.StatusCreated {
		t.Fatalf("expected 201, got %d", recorder.Code)
	}
    // 7) body tamamen boş mu kontrol et
	if recorder.Body.String() != "" {
		t.Fatalf("expected empty body, got %s", recorder.Body.String())
	}
}

func TestStatusHandler_InvalidCode(t *testing.T) {
    // 1) gin test mode
	gin.SetMode(gin.TestMode)
    // 2) router oluştur
	r := gin.Default()
	r.GET("/status/:code", StatusHandler)
    // 3) isteği oluştur
	req := httptest.NewRequest("GET", "/status/700", nil)
    // 4) recorder
	recorder := httptest.NewRecorder()
    // 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", recorder.Code)
	}
}

func TestStatusHandler_Teapot(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/status/:code", StatusHandler)

	req := httptest.NewRequest("GET", "/status/418", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusTeapot {
		t.Fatalf("expected 418, got %d", recorder.Code)
	}

	expectedBody := "I'm a teapot"
	if recorder.Body.String() != expectedBody {
		t.Fatalf("expected body: %q, got %q", expectedBody, recorder.Body.String())
	}
}
