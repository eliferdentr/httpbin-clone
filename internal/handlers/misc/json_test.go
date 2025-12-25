package misc

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

/*
TestJSONHandler

Kontroller:
1) Status code == 200
2) Content-Type == application/json
3) Body JSON parse edilebiliyor mu?
4) JSON içinde "slideshow" alanı var mı?
5) slideshow.author alanı var mı?
*/

func TestJSONHandler(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)

	// 2) router
	r := gin.Default()
	r.GET("/json", JSONHandler)

	// 3) request
	req := httptest.NewRequest("GET", "/json", nil)
	recorder := httptest.NewRecorder()

	// 4) serve
	r.ServeHTTP(recorder, req)

	// 5) status code
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}
	// 7) Content-Type kontrolü
	ct := recorder.Header().Get("Content-Type")
	if !strings.Contains(ct, "application/json") {
		t.Fatalf("expected application/json, got %s", ct)
	}

	// 8) JSON unmarshal
	var body map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &body)

	// 9) Alanlar var mı?
	if body["author_name"] == nil || body["author_surname"] == nil || body["author_github"] == nil {
		t.Fatalf("json response body missing")
	}

}
