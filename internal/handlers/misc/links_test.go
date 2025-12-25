package misc

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLinksHandler_Basic(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/links/:n", LinksHandler)

	req := httptest.NewRequest("GET", "/links/3", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	var body map[string]interface{}
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	links, ok := body["links"].([]interface{})
	if !ok {
		t.Fatalf("links field missing or invalid")
	}

	if len(links) != 3 {
		t.Fatalf("expected 3 links, got %d", len(links))
	}

	if links[0] != "http://example.com/links/0" {
		t.Fatalf("unexpected first link: %v", links[0])
	}
}

func TestLinksHandler_Invalid(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/links/:n", LinksHandler)

	req := httptest.NewRequest("GET", "/links/abc", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", recorder.Code)
	}
}
