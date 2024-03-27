package main

import (
	"github.com/Zeta-Manu/manu-lesson/config"
	"github.com/Zeta-Manu/manu-lesson/internal/application"
)

// @title Manu Swagger API
// @version 1.0
// @description server

// @host localhost:8080
// @BasePath /api
func main() {
	appConfig := config.NewAppConfig()

	application.NewApplication(*appConfig)
}
