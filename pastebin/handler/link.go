package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"main/database"
	"main/model"
	"time"
)

func LinkHandler(c *fiber.Ctx) error {
	fmt.Println()
	result := database.DB.Delete(&model.Text{}, "expire_at < ?",
		string(time.Now().Format("2006-01-02 15:04:05"))+".000000")
	fmt.Println(result)

	hash := c.Params("text_hash")
	var textRow model.Text
	database.DB.Model(&model.Text{}).Where("text_hash = ?", hash).First(&textRow)

	text := GetText(textRow.FileName)

	return c.Render("templates/link.html", fiber.Map{
		"text": text,
	})
}
