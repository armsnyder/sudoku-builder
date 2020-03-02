package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"net/http/httputil"
)

type configSpec struct {
	GinMode    string `default:"debug"`
	Static     bool   `default:"false"`
	StaticRoot string `default:"/var/www"`
}

func main() {
	config := new(configSpec)
	envconfig.MustProcess("sudoku", config)
	gin.SetMode(config.GinMode)
	router := gin.Default()
	if config.Static {
		router.NoRoute(static.ServeRoot("/", config.StaticRoot))
	} else {
		router.NoRoute(reverseProxyUI)
	}
	log.Fatal(router.Run(":8080"))
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
