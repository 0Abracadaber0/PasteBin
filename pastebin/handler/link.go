package handler

import "github.com/gofiber/fiber/v2"

func LinkHandler(c *fiber.Ctx) error {
	return c.Render("templates/link.html", fiber.Map{
		"hash": c.Params("text_hash"),
	})
}
