package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"main/database"
	"main/model"
)

func LinkHandler(c *fiber.Ctx) error {
	hash := c.Params("hash")
	result := database.DB.First(&model.Text{TextHash: hash})
	var textContent string
	if result.RowsAffected == 1 {
		result.Scan(&textContent)
	} else {
		// Обработайте случай, когда запись не найдена
	}
	fmt.Println(textContent)
	return c.Render("templates/link.html", fiber.Map{
		"hash": c.Params("text_hash"),
	})
}
