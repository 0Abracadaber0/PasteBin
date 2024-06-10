package handler

import "github.com/gofiber/fiber/v2"

func NotFoundHandler(c *fiber.Ctx) error {
	return c.SendFile("templates/404.html")
}
