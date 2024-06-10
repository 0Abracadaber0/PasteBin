package router

import (
	"github.com/gofiber/fiber/v2"
	"main/handler"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", handler.MainHandler)
	app.Post("/paste", handler.PasteHandler)

	text := app.Group("/text")
	text.Get("/:text_hash", handler.TextHandler)

	app.Use(handler.NotFoundHandler)
}
