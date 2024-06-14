package handler

import (
	"github.com/gofiber/fiber/v2"
	"main/generator"
)

func CalcHashHandler(c *fiber.Ctx) error {
	hash := generator.TakeNumber()
	return c.JSON(fiber.Map{
		"hash": hash,
	})
}
