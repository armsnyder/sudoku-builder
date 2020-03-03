package router

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
)

func New(isUIStatic bool, staticRoot string) *gin.Engine {
	engine := gin.Default()
	engine.Use(handleErrors)
	engine.POST("/solve", solve)
	if isUIStatic {
		engine.NoRoute(static.ServeRoot("/", staticRoot))
	} else {
		engine.NoRoute(reverseProxyUI)
	}
	return engine
}

func reverseProxyUI(c *gin.Context) {
	(&httputil.ReverseProxy{
		Director: func(req *http.Request) {
			log.Println(req)
			req.URL.Scheme = "http"
			req.URL.Host = "localhost:3000"
		},
	}).ServeHTTP(c.Writer, c.Request)
}
