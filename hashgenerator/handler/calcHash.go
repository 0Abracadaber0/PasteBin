package handler

import "github.com/gofiber/fiber/v2"

func CalcHashHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"hash": "HIIIII",
	})
}
