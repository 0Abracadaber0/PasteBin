package handler

import "github.com/gofiber/fiber/v2"

func PasteHandler(c *fiber.Ctx) error {
	return c.SendFile("templates/paste.html")
}
