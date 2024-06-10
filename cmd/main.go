package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", mainHandler)
	app.Post("/paste", pasteHandler)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendFile("templates/404.html")
	})

	log.Fatal(app.Listen(":8080"))
}

func mainHandler(c *fiber.Ctx) error {
	return c.SendFile("templates/index.html")
}

func pasteHandler(c *fiber.Ctx) error {
	return c.SendFile("templates/paste.html")
}
