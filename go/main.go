package main

import (
	"log"
	"sudoku-builder/router"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
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
	log.Fatal(router.New(config.Static, config.StaticRoot).Run(":8080"))
}
