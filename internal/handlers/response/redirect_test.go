package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRedirectHandler_SingleRedirect(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/redirect/:n", RedirectHandler)
	// 3) request
	req := httptest.NewRequest("GET", "/redirect/1", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	// 5) Status = 302
	if recorder.Code != http.StatusFound {
		t.Fatalf("expected %d, got %d", http.StatusFound, recorder.Code)
	}
	// 6) Location = "/redirect/0"
	if recorder.Header().Get("Location") != "/redirect/0" {
		t.Fatalf("expected /redirect/0, got %s", recorder.Header().Get("Location"))
	}
}

func TestRedirectHandler_FinalRedirectToGet(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/redirect/:n", RedirectHandler)
	// 3) request
	req := httptest.NewRequest("GET", "/redirect/0", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	// 5) Status = 302
	if recorder.Code != http.StatusFound {
		t.Fatalf("expected %d, got %d", http.StatusFound, recorder.Code)
	}
	// 6) Location = "/redirect/0"
	if recorder.Header().Get("Location") != "/get" {
		t.Fatalf("expected /get, got %s", recorder.Header().Get("Location"))
	}
}

func TestRedirectHandler_InvalidN(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/redirect/:n", RedirectHandler)
	// 3) request
	req := httptest.NewRequest("GET", "/redirect/abc", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	// 5) Status = 302
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestRelativeRedirectHandler_Chain(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/relative-redirect/:n", RelativeRedirectHandler)
	// 3) request
	req := httptest.NewRequest("GET", "/relative-redirect/2", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	// 5) Location = "/relaitve-redirect/0"
	if recorder.Header().Get("Location") != "/relative-redirect/1" {
		t.Fatalf("expected /relative-redirect/1, got %s", recorder.Header().Get("Location"))
	}
}

func TestRelativeRedirectHandler_ToGet(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/relative-redirect/:n", RelativeRedirectHandler)
	// 3) request
	req := httptest.NewRequest("GET", "/relative-redirect/0", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	// 5) Location = "/relaitve-redirect/0"
	if recorder.Header().Get("Location") != "/get" {
		t.Fatalf("expected http://example.com/get, got %s", recorder.Header().Get("Location"))
	}
}

func TestRelativeRedirectHandler_Invalid(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/relative-redirect/:n", RelativeRedirectHandler)
	// 3) request
	req := httptest.NewRequest("GET", "/relative-redirect/xyz", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}
func TestAbsoluteRedirectHandler_Chain(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/absolute-redirect/:n", AbsoluteRedirectHandler)
	// 3) request
	req := httptest.NewRequest("GET", "/absolute-redirect/5", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	if recorder.Header().Get("Location") != "http://example.com/absolute-redirect/4" {
		t.Fatalf("expected http://example.com/absolute-redirect/4, got %s", recorder.Header().Get("Location"))
	}
}

func TestAbsoluteRedirectHandler_ToGet(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/absolute-redirect/:n", AbsoluteRedirectHandler)
	// 3) request
	req := httptest.NewRequest("GET", "/absolute-redirect/0", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	if recorder.Header().Get("Location") != "http://example.com/get" {
		t.Fatalf("expected http://example.com/get, got %s", recorder.Header().Get("Location"))
	}
}

func TestAbsoluteRedirectHandler_Invalid(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router oluştur
	r := gin.Default()
	r.GET("/absolute-redirect/:n", AbsoluteRedirectHandler)
	// 3) request
	req := httptest.NewRequest("GET", "/absolute-redirect/asd", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}
