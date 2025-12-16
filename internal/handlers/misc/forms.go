package misc

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
FormsHandler

HttpBin /forms/post davranışı:

1) Endpoint:
   POST /forms/post

2) Kabul ettiği Content-Type’lar:
   - application/x-www-form-urlencoded
   - multipart/form-data

3) Request’ten okunacaklar:
   - Form field’lar (key=value)
   - Eğer multipart ise:
       - dosyalar (file)
       - form alanları

4) Response JSON formatı:
{
  "form": {              // text alanları
    "field1": "value1",
    "field2": "value2"
  },
  "files": {             // dosyalar (varsa)
    "file": "filename.txt"
  }
}

5) Eğer form boşsa:
   - form = {}
   - files = {}

6) Status code:
   - Her zaman 200 OK
   - Validation yok (HttpBin gibi)

7) Header / method fark etmez:
   - POST beklenir ama strict davranılmaz
*/

func FormsPostHandler(c *gin.Context) {
	form := map[string]string{}
	files := map[string]string{}

	ct := c.GetHeader("Content-Type")

	if strings.HasPrefix(ct, "application/x-www-form-urlencoded") {
		_ = c.Request.ParseForm()
		for k, v := range c.Request.PostForm {
			form[k] = v[0]
		}
	}

	if strings.HasPrefix(ct, "multipart/form-data") {
		mf, err := c.MultipartForm()
		if err == nil {
			for k, v := range mf.Value {
				form[k] = v[0]
			}
			for k, v := range mf.File {
				files[k] = v[0].Filename
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"form":  form,
		"files": files,
	})
}

func getMultipartFormData(c *gin.Context) (map[string][]string, map[string]string, error) {
	form := map[string][]string{}
	files := map[string]string{}

	mf, err := c.MultipartForm()
	if err != nil {
		return form, files, err
	}

	for k, v := range mf.Value {
		form[k] = v
	}

	for k, f := range mf.File {
		if len(f) > 0 {
			files[k] = f[0].Filename
		}
	}

	return form, files, nil
}

func getFormFields(c *gin.Context) (map[string][]string, error) {
	if err := c.Request.ParseForm(); err != nil {
		return map[string][]string{}, err
	}
	return c.Request.PostForm, nil
}
