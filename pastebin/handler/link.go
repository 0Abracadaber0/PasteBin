package handler

import (
	"github.com/gofiber/fiber/v2"
	"main/database"
	"main/model"
)

func LinkHandler(c *fiber.Ctx) error {
	hash := c.Params("text_hash")
	var textRow model.Text
	database.DB.Model(&model.Text{}).Where("text_hash = ?", hash).First(&textRow)

	text := GetText(textRow.FileName)

	return c.Render("templates/link.html", fiber.Map{
		"text": text,
	})
}
