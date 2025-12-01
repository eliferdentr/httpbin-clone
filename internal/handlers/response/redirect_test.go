package response

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
)

func TestRedirectHandler_Basic(t *testing.T) {
	// 1) gin test mode
	gin.SetMode(gin.TestMode)
	// 2) router
	r := gin.Default()
	r.GET("/range/:n", RangeHandler)
	// 3) request: /range/5
	req := httptest.NewRequest("GET", "/range/5", nil)
	// 4) recorder
	recorder := httptest.NewRecorder()
	// 5) ServeHTTP
	r.ServeHTTP(recorder, req)
	// 6) status = 200
	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected code: %d . Received: %d ", http.StatusOK, recorder.Code)
	}
	// 7) Content-Type = application/octet-stream
	expectedContentType := "application/octet-stream"
	contentTyoe := recorder.Header().Get("Content-Type")
	if contentTyoe == "" || contentTyoe != expectedContentType {
		t.Fatalf("Expected Content-Type: %q . Received: %q ", expectedContentType, contentTyoe)
	}
	// 8) Body length = 5 byte
	bodyLength := recorder.Body.Len()
	expectedBodyLength := 5
	if bodyLength != expectedBodyLength {
		t.Fatalf("Expected bodyLength: %d . Received: %d ", expectedBodyLength, bodyLength)
	}

}
