package response

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestRangeHandler_Basic(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/range/:n", RangeHandler)
	// 3) request: /range/5
	req := httptest.NewRequest("GET", "/range/5", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) status = 200
	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected code: %d . Received: %d ", http.StatusOK, recorder.Code)
	}
	// 7) Content-Type = application/octet-stream
	expectedContentType := "application/octet-stream"
	contentTyoe := recorder.Header().Get("Content-Type")
	if contentTyoe == "" || contentTyoe != expectedContentType {
		t.Fatalf("Expected Content-Type: %q . Received: %q ", expectedContentType, contentTyoe)
	}
	// 8) Body length = 5 byte
	bodyLength := recorder.Body.Len()
	expectedBodyLength := 5
	if bodyLength != expectedBodyLength {
		t.Fatalf("Expected bodyLength: %d . Received: %d ", expectedBodyLength, bodyLength)
	}

}

func TestRangeHandler_InvalidNumber(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/range/:n", RangeHandler)
	// 3) "/range/abc" → 400
	req := httptest.NewRequest("GET", "/range/abc", nil)
	recorder := httptest.NewRecorder()
	// 4) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 8) status = 400
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected code: %d . Received: %d ", http.StatusBadRequest, recorder.Code)
	}
}

func TestRangeHandler_NegativeNumber(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/range/:n", RangeHandler)
	// 3) "/range/-50" → 400
	req := httptest.NewRequest("GET", "/range/-50", nil)
	recorder := httptest.NewRecorder()
	// 4) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 8) status = 400
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("Expected code: %d . Received: %d ", http.StatusBadRequest, recorder.Code)
	}
}

func TestRangeHandler_DurationStreaming(t *testing.T) {
    gin.SetMode(gin.TestMode)

    r := gin.Default()
    r.GET("/range/:n", RangeHandler)

    req := httptest.NewRequest("GET", "/range/5?duration=1", nil)
    recorder := httptest.NewRecorder()

    start := time.Now()
    r.ServeHTTP(recorder, req)
    elapsed := time.Since(start)

    if recorder.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", recorder.Code)
    }

    if elapsed < time.Second {
        t.Fatalf("expected at least 1 second duration, got %v", elapsed)
    }

    if recorder.Body.Len() != 5 {
        t.Fatalf("expected 5 bytes, got %d", recorder.Body.Len())
    }
}

func TestRangeHandler_ZeroBytes(t *testing.T) {
    gin.SetMode(gin.TestMode)

    r := gin.Default()
    r.GET("/range/:n", RangeHandler)

    req := httptest.NewRequest("GET", "/range/0", nil)
    recorder := httptest.NewRecorder()

    start := time.Now()
    r.ServeHTTP(recorder, req)
    elapsed := time.Since(start)

    if recorder.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", recorder.Code)
    }

    if recorder.Body.Len() != 0 {
        t.Fatalf("expected 0 bytes, got %d", recorder.Body.Len())
    }

    if elapsed > 200*time.Millisecond {
        t.Fatalf("range/0 should finish immediately, took %v", elapsed)
    }
}
