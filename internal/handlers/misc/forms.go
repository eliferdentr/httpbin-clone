package misc

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
POST /forms/post

HttpBin uyumlu davranış:

Response:
{
  "form":  { "field": "value" },
  "files": { "file": "filename.txt" }
}

- application/x-www-form-urlencoded destekler
- multipart/form-data destekler
- Dosya content'i dönmez, SADECE filename döner
- Form / file yoksa boş map döner
- Status her zaman 200
*/

// FormsPostHandler godoc
//
// @Summary      Submit form data
// @Description Accepts form-urlencoded or multipart form data
// @Tags         misc
// @Accept       application/x-www-form-urlencoded
// @Accept       multipart/form-data
// @Produce      application/json
// @Success      200 {object} map[string]interface{}
// @Router       /forms/post [post]
func FormsPostHandler(c *gin.Context) {
	form := map[string]string{}
	files := map[string]string{}

	ct := c.GetHeader("Content-Type")

	// x-www-form-urlencoded
	if strings.HasPrefix(ct, "application/x-www-form-urlencoded") {
		_ = c.Request.ParseForm()
		for k, v := range c.Request.PostForm {
			if len(v) > 0 {
				form[k] = v[0]
			}
		}
	}

	// multipart/form-data
	if strings.HasPrefix(ct, "multipart/form-data") {
		mf, err := c.MultipartForm()
		if err == nil {
			for k, v := range mf.Value {
				if len(v) > 0 {
					form[k] = v[0]
				}
			}
			for k, v := range mf.File {
				if len(v) > 0 {
					files[k] = v[0].Filename
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"form":  form,
		"files": files,
	})
}
