package request

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAnythingHandler_Get(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// 3) GET /anything?x=5 isteği oluştur
	r.GET("/anything", AnythingHandler)
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/anything?x=5", nil)
	r.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}
	// JSON parse
	var data map[string]interface{}
	if err := json.Unmarshal(recorder.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	if data["method"] != "GET" {
		t.Fatalf("expected method GET, got %v", data["method"])
	}

	query := data["query"].(map[string]interface{})
	if query["x"].([]interface{})[0] != "5" {
		t.Fatalf("expected query x=5, got %v", query["x"])
	}

	if data["url"] != "/anything?x=5" {
		t.Fatalf("expected url /anything?x=5, got %v", data["url"])
	}
}

func TestAnythingHandler_PostBody(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/anything", AnythingHandler)

	body := `{"x":10}`

	req := httptest.NewRequest("POST", "/anything", strings.NewReader(body))
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	var data map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &data)

	if data["body"] != body {
		t.Fatalf("expected body=%s, got %v", body, data["body"])
	}
}

func TestAnythingHandler_Headers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/anything", AnythingHandler)

	req := httptest.NewRequest("GET", "/anything", nil)
	req.Header.Set("X-Test", "123")

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	var data map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &data)

	headers := data["headers"].(map[string]interface{})

	if headers["X-Test"].([]interface{})[0] != "123" {
		t.Fatalf("expected header X-Test=123, got %v", headers["X-Test"])
	}
}
