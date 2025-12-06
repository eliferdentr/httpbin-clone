package request

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetHandler_Get(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// 3) GET /get?x=5 isteği oluştur
	r.GET("/get", GetHandler)
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/get?x=10&y=yes", nil)
	r.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	// JSON parse
	var data map[string]interface{}
	if err := json.Unmarshal(recorder.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	// args içinde x=10 var mı?
	query := data["args"].(map[string]interface{})
	if query["x"].([]interface{})[0] != "10" {
		t.Fatalf("expected query x=10, got %v", query["x"])
	}
	// args içinde y=yes var mı?
	if query["y"].([]interface{})[0] != "yes" {
		t.Fatalf("expected query y=yes, got %v", query["y"])
	}
	// headers döndü mü?
	if data["headers"] == "" {
		t.Fatal("expected headers")
	}
	// origin boş değil mi?
	if data["origin"] == "" {
		t.Fatal("expected origin")
	}
	// url doğru mu?
	if data["url"] != "/get?x=10&y=yes" {
		t.Fatalf("expected url get?x=10&y=yes, got %v", data["url"])
	}
}
