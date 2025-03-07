package main

import (
	"github.com/om13rajpal/gobackend/config"
	"github.com/om13rajpal/gobackend/internal/routes"
)

func main() {
	config.LoadConfig()

	router := routes.InitRouter()
	router.Run(":" + config.PORT)
}