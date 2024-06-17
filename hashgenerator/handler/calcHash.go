package handler

import (
	"github.com/gofiber/fiber/v2"
	"main/generator"
)

func CalcHashHandler(c *fiber.Ctx) error {
	hash := generator.TakeHash()

	return c.JSON(fiber.Map{
		"hash": hash,
	})
}
