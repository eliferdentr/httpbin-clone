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

	body := "name=kitten&age=5"
	req := httptest.NewRequest(
		"POST",
		"/forms/post",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	form := resp["form"].(map[string]interface{})

	if form["name"].([]interface{})[0] != "kitten" {
		t.Fatalf("expected name=kitten")
	}
	if form["age"].([]interface{})[0] != "5" {
		t.Fatalf("expected age=5")
	}
}

func TestFormsPostHandler_Multipart(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/forms/post", FormsPostHandler)

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// form field
	_ = writer.WriteField("username", "kitten")

	// fake file
	fileWriter, _ := writer.CreateFormFile("file", "test.txt")
	fileWriter.Write([]byte("hello"))

	writer.Close()

	req := httptest.NewRequest("POST", "/forms/post", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid json: %v", err)
	}

	files := resp["files"].(map[string]interface{})

	if files["file"] != "test.txt" {
		t.Fatalf("expected file test.txt, got %v", files["file"])
	}
	
	if len(files) == 0 {
		t.Fatalf("expected files")
	}
}

func TestFormsPostHandler_UnsupportedContentType(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/forms/post", FormsPostHandler)

	req := httptest.NewRequest(
		"POST",
		"/forms/post",
		strings.NewReader(`{"x":1}`),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}
