package misc

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	// testing
	// net/http
	// net/http/httptest
	// gin
	"testing"

	"github.com/gin-gonic/gin"
)

/*
TestDenyHandler

Kontroller:
1) Status code == 403 mü?
2) Body JSON mu?
3) message alanı beklenen string mi?
*/

func TestDenyHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/deny", DenyHandler)
	req := httptest.NewRequest("GET", "/deny", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	// 6) status code kontrolü
	if recorder.Code != http.StatusForbidden {
		t.Errorf("DenyHandler should return HTTP 403 status, got %d", recorder.Code)
	}

	// 7) body JSON parse
	var data map[string]interface{}
	if err := json.Unmarshal(recorder.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	if data["message"] != "YOU SHALL NOT PASS" {
		t.Errorf("DenyHandler should return 'YOU SHALL NOT PASS', got %s", data["message"])
	}
}
