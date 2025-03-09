package main

import (
	"github.com/om13rajpal/gobackend/config"
	"github.com/om13rajpal/gobackend/internal/database"
	"github.com/om13rajpal/gobackend/internal/routes"
)

func main() {
	config.LoadConfig()

	database.ConnectMongo()
	database.ConnectPostgres()
	defer database.Pool.Close()

	router := routes.InitRouter()
	router.Run(":" + config.PORT)
}