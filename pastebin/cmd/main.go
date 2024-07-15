package main

import (
	"github.com/gofiber/fiber/v2"
	"main/config"
	"main/database"
	"main/router"
)

func main() {
	log := config.SetupLogger(config.Config("ENV"))
	log.Info("starting server", "env", config.Config("ENV"))

	app := fiber.New()

	database.ConnectDB()

	router.SetUpRoutes(app)

	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		return
	}
}
