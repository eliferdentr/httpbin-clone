package router

import (
	"github.com/gin-gonic/gin"

	// auth
	"httbinclone-eliferden.com/internal/handlers/auth"

	// response
	"httbinclone-eliferden.com/internal/handlers/response"

	// request
	"httbinclone-eliferden.com/internal/handlers/request"

	// cookies
	"httbinclone-eliferden.com/internal/handlers/cookies"

	// compression
	"httbinclone-eliferden.com/internal/handlers/compression"

	// images
	"httbinclone-eliferden.com/internal/handlers/images"

	// misc
	"httbinclone-eliferden.com/internal/handlers/misc"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "httbinclone-eliferden.com/docs" // <-- module path senin go.mod’daki module adı olmalı
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//==========================
	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// =========================
	// AUTH
	// =========================
	r.GET("/basic-auth/:user/:passwd", auth.BasicAuthHandler)
	r.GET("/bearer", auth.BearerAuthHandler)
	r.GET("/digest-auth/:qop/:user/:passwd", auth.DigestAuthHandler)
	r.GET("/hidden-basic-auth/:user/:passwd", auth.HiddenBasicAuthHandler)

	// =========================
	// RESPONSE
	// =========================
	r.GET("/status/:code", response.StatusHandler)
	r.GET("/bytes/:n", response.BytesHandler)
	r.GET("/cache/:n", response.CacheHandler)
	r.GET("/delay/:n", response.DelayHandler)
	r.GET("/drip", response.DripHandler)
	r.GET("/range/:n", response.RangeHandler)
	r.GET("/redirect/:n", response.RedirectHandler)
	r.GET("/relative-redirect/:n", response.RelativeRedirectHandler)
	r.GET("/absolute-redirect/:n", response.AbsoluteRedirectHandler)
	r.GET("/stream/:n", response.StreamHandler)

	// =========================
	// REQUEST
	// =========================
	r.Any("/anything", request.AnythingHandler)
	r.GET("/get", request.GetHandler)
	r.GET("/headers", request.HeadersHandler)
	r.GET("/ip", request.IPHandler)
	r.Any("/methods", request.MethodsHandler)
	r.GET("/user-agent", request.UserAgentHandler)

	// =========================
	// COOKIES
	// =========================
	r.GET("/cookies", cookies.CookiesHandler)
	r.GET("/cookies/set/:name/:value", cookies.SetCookieHandler)
	r.GET("/cookies/delete", cookies.DeleteCookieHandler)

	// =========================
	// COMPRESSION
	// =========================
	r.GET("/gzip", compression.GzipHandler)
	r.GET("/brotli", compression.BrotliHandler)
	r.GET("/deflate", compression.DeflateHandler)

	// =========================
	// IMAGES
	// =========================
	r.GET("/image/jpeg", images.JpegHandler)
	r.GET("/image/png", images.PNGHandler)
	r.GET("/image/svg", images.SVGHandler)
	r.GET("/image/webp", images.WEBPHandler)

	// =========================
	// MISC
	// =========================
	r.GET("/deny", misc.DenyHandler)
	r.POST("/forms/post", misc.FormsPostHandler)
	r.GET("/html", misc.HTMLHandler)
	r.GET("/json", misc.JSONHandler)
	r.GET("/links/:n", misc.LinksHandler)
	r.GET("/robots.txt", misc.RobotsHandler)
	r.GET("/uuid", misc.UUIDHandler)
	r.GET("/xml", misc.XMLHandler)

	return r
}
