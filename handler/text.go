package handler

import "github.com/gofiber/fiber/v2"

func TextHandler(c *fiber.Ctx) error {
	return c.Render("templates/text.html", fiber.Map{
		"text": c.Params("text_hash"),
	})
}
