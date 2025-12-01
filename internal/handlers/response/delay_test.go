package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/gin-gonic/gin"
)

func TestDelayHandler_Valid(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/delay/:n", DelayHandler)
	// 3) request: /delay/1
	req := httptest.NewRequest("GET", "/delay/1", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) start := time.Now()
	start := time.Now()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 7) duration := time.Since(start)
	duration := time.Since(start)
	// 8) status = 200
	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected code: %d . Received: %d ", http.StatusOK, recorder.Code)
	}
	// 9) duration >= 1 second olmalı
	if duration < 1 {
		t.Fatalf("Duration is not smaller than 1. Duration: %s", duration)
	}
	// 10) body içinde {"delay":1} var mı?
	var result map[string]int

	if err := json.Unmarshal(recorder.Body.Bytes(), &result); err != nil {
		t.Fatalf("invalid JSON returned: %v", err)
	}

	if result["delay"] != 1 {
		t.Fatalf("expected delay=1, got %d", result["delay"])
	}

}

func TestDelayHandler_InvalidNumber(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/delay/:n", DelayHandler)
	// 3) "/delay/abc" → 400
	req := httptest.NewRequest("GET", "/delay/abc", nil)
	recorder := httptest.NewRecorder()
	// 4) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 8) status = 400
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected code: %d . Received: %d ", http.StatusBadRequest, recorder.Code)
	}
}

func TestDelayHandler_NegativeNumber(t *testing.T) {
		// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/delay/:n", DelayHandler)
	// 3) "/delay/-50" → 400
	req := httptest.NewRequest("GET", "/delay/-50", nil)
	recorder := httptest.NewRecorder()
	// 4) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 8) status = 400
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected code: %d . Received: %d ", http.StatusBadRequest, recorder.Code)
	}
}
