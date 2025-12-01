package response

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestDripHandler_Basic(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/drip", DripHandler)
	// 3) request: /drip?numbytes=5&duration=1&delay=0
	req := httptest.NewRequest("GET", "/drip?numbytes=5&duration=1&delay=0", nil)
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
	if duration < time.Second {
		t.Fatalf("Duration is smaller than 1. Duration: %s", duration)
	}
	// 10) Body length = 5 byte?
	if recorder.Body.Len() != 5 {
		t.Fatalf("Expected body length is 5. Received: %d", recorder.Body.Len())
	}

}

func TestDripHandler_DelayOnly(t *testing.T) {
	//
	// 
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/drip", DripHandler)
	// 3) request:  /drip?numbytes=1&duration=0&delay=2  
	req := httptest.NewRequest("GET", "/drip?numbytes=1&duration=0&delay=2", nil)
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
	// 9) response en az 2 saniye sürmeli
	if duration < time.Second * 2 {
		t.Fatalf("Duration is smaller than 2. Duration: %s", duration)
	}
}

func TestDripHandler_InvalidNumBytes(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/drip", DripHandler)
	// 3) request:  
	req := httptest.NewRequest("GET", "/drip?numbytes=abc&duration=0&delay=2", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	// 5) status = 200
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected code: %d . Received: %d ", http.StatusBadRequest, recorder.Code)
	}
}

func TestDripHandler_InvalidDuration(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/drip", DripHandler)
	// 3) request:  
	req := httptest.NewRequest("GET", "/drip?numbytes=1&duration=-1&delay=2", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	// 5) status = 200
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected code: %d . Received: %d ", http.StatusBadRequest, recorder.Code)
	}
}

func TestDripHandler_InvalidDelay(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/drip", DripHandler)
	// 3) request:  
	req := httptest.NewRequest("GET", "/drip?numbytes=1&duration=1&delay=-5", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	// 5) status = 200
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected code: %d . Received: %d ", http.StatusBadRequest, recorder.Code)
	}
}

func TestDripHandler_InvalidCode(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/drip", DripHandler)
	// 3) request:  
	req := httptest.NewRequest("GET", "/drip?numbytes=1&duration=1&delay=5&code=600", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	// 5) status = 200
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected code: %d . Received: %d ", http.StatusBadRequest, recorder.Code)
	}
}
