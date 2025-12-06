package request

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUserAgentHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/user-agent", UserAgentHandler)

	req := httptest.NewRequest("GET", "/user-agent", nil)
	req.Header.Set("User-Agent", "KittenBrowser/1.0")

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(recorder.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	if data["user-agent"] != "KittenBrowser/1.0" {
		t.Fatalf("expected user-agent KittenBrowser/1.0, got %v", data["user-agent"])
	}
}
