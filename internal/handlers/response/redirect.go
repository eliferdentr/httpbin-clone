package response

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/utils"
)

func RedirectHandler(c *gin.Context) {
	n, invalid := utils.GetNParam(c)
	if invalid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid redirect count"})
		return
	}

	if n == 1 {
		c.Redirect(http.StatusFound, "/get")
	} else {
		c.Redirect(http.StatusFound, "/redirect/"+strconv.Itoa(n-1))
	}
}

func RedirectToHandler(c *gin.Context) {
	target := c.Query("url")
	if target == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url parameter is required"})
		return
	}

	status := http.StatusFound
	if codeStr := c.Query("status_code"); codeStr != "" {
		code, err := strconv.Atoi(codeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status code"})
			return
		}
		switch code {
		case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther,
			http.StatusTemporaryRedirect, http.StatusPermanentRedirect:
			status = code
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "status code must be a valid redirect code"})
			return
		}
	}

	c.Redirect(status, target)
}

func AbsoluteRedirectHandler(c *gin.Context) {
	n, invalid := utils.GetNParam(c)
	if invalid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid redirect count"})
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	host := c.Request.Host

	if n == 1 {
		c.Redirect(http.StatusFound, scheme+"://"+host+"/get")
	} else {
		nextURL := scheme + "://" + host + "/absolute-redirect/" + strconv.Itoa(n-1)
		c.Redirect(http.StatusFound, nextURL)
	}
}


