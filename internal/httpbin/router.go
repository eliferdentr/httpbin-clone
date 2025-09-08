package httpbin

import (
	"github.com/gin-gonic/gin"
	"httbinclone-eliferden.com/internal/handlers/auth"
	"httbinclone-eliferden.com/internal/handlers/compression"
	"httbinclone-eliferden.com/internal/handlers/cookies"
	"httbinclone-eliferden.com/internal/handlers/images"
	"httbinclone-eliferden.com/internal/handlers/misc"
	"httbinclone-eliferden.com/internal/handlers/request"
	"httbinclone-eliferden.com/internal/handlers/response"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//request
	r.GET("/ip", request.IPHandler)
	r.GET("/user-agent", request.UserAgentHandler)
	r.GET("/headers", request.HeadersHandler)
	r.GET("/get", request.GetHandler)
	r.Any("/anything", request.AnythingHandler)


	//response
	r.GET("/status/:code", response.StatusHandler)
	r.GET("/redirect/:n", response.RedirectHandler)
	r.GET("/redirect-to", response.RedirectToHandler)
	r.GET("/absolute-redirect/:n", response.AbsoluteRedirectHandler)
	r.GET("/delay/:seconds", response.DelayHandler)
	r.GET("/drip", response.DripHandler)
	r.GET("/stream/:n", response.StreamHandler)
	r.GET("/bytes/:n", response.BytesHandler)
	r.GET("/range/:n", response.RangeHandler)

	//cookies
	r.GET("/cookies", cookies.GetCookiesHandler)
	r.GET("/cookies/set", cookies.SetCookiesHandler)
	r.GET("/cookies/delete", cookies.DeleteCookiesHandler)

	//auth
	r.GET("/basic-auth/:user/:passwd", auth.BasicAuthHandler)
	r.GET("/hidden-basic-auth/:user/:passwd", auth.HiddenBasicAuthHandler)
	r.GET("/digest-auth/:qop/:user/:passwd", auth.DigestAuthHandler)
	r.GET("/bearer", auth.BearerHandler)

	// compression
	r.GET("/gzip", compression.GzipHandler)
	r.GET("/deflate", compression.DeflateHandler)
	r.GET("/brotli", compression.BrotliHandler)

	// images
	r.GET("/image/png", images.PNGHandler)
	r.GET("/image/jpeg", images.JPEGHandler)
	r.GET("/image/svg", images.SVGHandler)
	r.GET("/image/webp", images.WEBPHandler)

	//misc
	r.GET("/html", misc.HTMLHandler)
	r.GET("/json", misc.JSONHandler)
	r.GET("/xml", misc.XMLHandler)
	r.GET("/robots.txt", misc.RobotsHandler)
	r.GET("/deny", misc.DenyHandler)
	r.GET("/links/:n", misc.LinksHandler)
	r.POST("/forms/post", misc.FormsHandler)
	r.GET("/uuid", misc.UUIDHandler)

	return r
}