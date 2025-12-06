package request

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHeadersHandler_ReturnsHeaders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/headers", HeadersHandler)

	req := httptest.NewRequest("GET", "/headers", nil)

	// custom header
	req.Header.Set("X-Test", "123")
	req.Header.Set("User-Agent", "MyAgent")

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(recorder.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	// headers map'ini al
	headers := data["headers"].(map[string]interface{})

	// X-Test doğru dönmüş mü?
	if headers["X-Test"].([]interface{})[0] != "123" {
		t.Fatalf("expected X-Test=123, got %v", headers["X-Test"])
	}

	// User-Agent doğru mu?
	if headers["User-Agent"].([]interface{})[0] != "MyAgent" {
		t.Fatalf("expected User-Agent=MyAgent, got %v", headers["User-Agent"])
	}
}
