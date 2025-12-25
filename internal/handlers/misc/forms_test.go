package misc

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFormsPostHandler_URLEncoded(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/forms/post", FormsPostHandler)

	body := strings.NewReader("name=kitten&age=3")
	req := httptest.NewRequest("POST", "/forms/post", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var resp map[string]interface{}
	_ = json.Unmarshal(rec.Body.Bytes(), &resp)

	form := resp["form"].(map[string]interface{})
	files := resp["files"].(map[string]interface{})

	if form["name"] != "kitten" {
		t.Fatalf("expected name=kitten")
	}
	if form["age"] != "3" {
		t.Fatalf("expected age=3")
	}
	if len(files) != 0 {
		t.Fatalf("expected no files")
	}
}

func TestFormsPostHandler_Multipart(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/forms/post", FormsPostHandler)

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	_ = writer.WriteField("username", "kitten")

	part, _ := writer.CreateFormFile("file", "test.txt")
	part.Write([]byte("hello"))

	writer.Close()

	req := httptest.NewRequest("POST", "/forms/post", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var resp map[string]interface{}
	_ = json.Unmarshal(rec.Body.Bytes(), &resp)

	form := resp["form"].(map[string]interface{})
	files := resp["files"].(map[string]interface{})

	if form["username"] != "kitten" {
		t.Fatalf("expected username=kitten")
	}

	if files["file"] != "test.txt" {
		t.Fatalf("expected file=test.txt")
	}
}

func TestFormsPostHandler_Empty(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/forms/post", FormsPostHandler)

	req := httptest.NewRequest("POST", "/forms/post", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	var resp map[string]interface{}
	_ = json.Unmarshal(rec.Body.Bytes(), &resp)

	form := resp["form"].(map[string]interface{})
	files := resp["files"].(map[string]interface{})

	if len(form) != 0 || len(files) != 0 {
		t.Fatalf("expected empty form and files")
	}
}
