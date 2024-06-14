package router

import (
	"github.com/gofiber/fiber/v2"
	"main/handler"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/calc_hash", handler.CalcHashHandler)
}
