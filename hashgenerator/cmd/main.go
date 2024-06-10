package main

import (
	"github.com/gofiber/fiber/v2"
	"hashgenerator/router"
	"log"
)

func main() {
	app := fiber.New()

	router.SetUpRoutes(app)

	log.Fatal(app.Listen(":8081"))
}