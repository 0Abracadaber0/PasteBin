package handler

import (
	"github.com/gofiber/fiber/v2"
	"main/generator"
)

type Ticket struct {
	Number int `json:"number"`
}

func CalcHashHandler(c *fiber.Ctx) error {
	var ticket Ticket

	if err := c.BodyParser(&ticket); err != nil {
		return err
	}

	hash := generator.ConvertNumber(ticket.Number)

	return c.JSON(fiber.Map{
		"hash": hash,
	})
}
