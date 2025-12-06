package request

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMethodsHandler_GET(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/methods", MethodsHandler)

	req := httptest.NewRequest("GET", "/methods?x=10", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	var data map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &data)
	if err != nil {
		t.Fatal(err)
	}

	if data["method"] != "GET" {
		t.Fatalf("expected GET, got %v", data["method"])
	}
}

func TestMethodsHandler_POST(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/methods", MethodsHandler)

	body := `{"a":1}`
	req := httptest.NewRequest("POST", "/methods", strings.NewReader(body))
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	var data map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &data)
	if err != nil {
		t.Fatal(err)
	}

	// check method
	if data["method"] != "POST" {
		t.Fatalf("expected POST, got %v", data["method"])
	}

	// check raw body
	if data["body"] != body {
		t.Fatalf("expected body %s, got %v", body, data["body"])
	}
}
