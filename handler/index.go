package handler

import "github.com/gofiber/fiber/v2"

func MainHandler(c *fiber.Ctx) error {
	return c.SendFile("templates/index.html")
}
