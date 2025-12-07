package cookies

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// /cookies endpoint → mevcut cookie'leri döndürüyor mu?
func TestCookiesHandler_ReturnsCookies(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/cookies", CookiesHandler)
	// 3) request: Cookie header ekle
	req := httptest.NewRequest("GET", "/cookies", nil)
	req.AddCookie(&http.Cookie{
		Name:     "x",
		Value:    "10",
		MaxAge:   1,
		HttpOnly: false,
		Path:     "/",
	})
	// 4) ServeHTTP
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	// 5) JSON parse → cookies map doğru mu?
	var data map[string]interface{}
	if err := json.Unmarshal(recorder.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	// cookies alanı var mı?
	cookiesField := data["cookies"]
	if cookiesField == nil {
		t.Fatalf("cookies field missing")
	}

	// map'e cast et
	cookies := cookiesField.(map[string]interface{})

	// x cookie’si var mı?
	if cookies["x"] != "10" {
		t.Fatalf("expected cookie x=10, got %v", cookies["x"])
	}
}

// /cookies/set/:name/:value → cookie set edip redirect ediyor mu?
func TestSetCookieHandler_SetsCookie(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/cookies/set/:name/:value", SetCookieHandler)
	// 3) request: Cookie header ekle
	req := httptest.NewRequest("GET", "/cookies/set/test/123", nil)
	// 4) ServeHTTP
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	// 5) Status = 302
	if recorder.Code != http.StatusFound {
		t.Fatalf("expected status code %d, got %d", http.StatusFound, recorder.Code)
	}
	// 6) Set-Cookie header doğru mu?
	if recorder.Header().Get("Set-Cookie") != "test=123; Path=/" {
		t.Fatalf("expected Set-Cookie %s, got %s", "test=123; Path=/", recorder.Header().Get("Set-Cookie"))
	}
	// 7) Location header /cookies mi?
	if recorder.Header().Get("Location") != "/cookies" {
		t.Fatalf("expected Location %s, got %s", "/cookies", recorder.Header().Get("Location"))
	}
}

// /cookies/delete?name=X → cookie silme çalışıyor mu?
func TestDeleteCookieHandler_DeletesCookie(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/cookies/delete", DeleteCookieHandler)

	req := httptest.NewRequest("GET", "/cookies/delete?name=test", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusFound {
		t.Fatalf("expected 302, got %d", recorder.Code)
	}

	// Set-Cookie header içinde Max-Age=-1 olmalı
	setCookie := recorder.Header().Get("Set-Cookie")
	//ginin default davranışı, max-age'i 0lıyor
	if !strings.Contains(setCookie, "Max-Age=0") {
		t.Fatalf("expected deleted cookie (Max-Age=0), got: %s", setCookie)
	}

	if recorder.Header().Get("Location") != "/cookies" {
		t.Fatalf("expected redirect to /cookies, got %s", recorder.Header().Get("Location"))
	}
}
