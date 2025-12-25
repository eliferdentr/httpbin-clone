package misc

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TestUUIDHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/uuid", UUIDHandler)

	req := httptest.NewRequest("GET", "/uuid", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	// Status code
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	// JSON parse
	var data map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	// uuid alanı var mı?
	uuidStr, ok := data["uuid"]
	if !ok || uuidStr == "" {
		t.Fatal("uuid field missing or empty")
	}

	// UUID geçerli mi?
	if _, err := uuid.Parse(uuidStr); err != nil {
		t.Fatalf("invalid uuid format: %v", err)
	}
}
