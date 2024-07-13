package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"main/database"
	"main/internal/converter"
	"main/model"
)

func PasteHandler(c *fiber.Ctx) error {
	url := "http://hashgenerator:8081/api/calc_hash"

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)

	if err := a.Parse(); err != nil {
		panic(err)
	}

	_, body, errs := a.Bytes()
	if errs != nil {
		log.Fatal(errs)
	}

	var hash fiber.Map
	err := json.Unmarshal(body, &hash)
	if err != nil {
		log.Fatal(err)
	}

	text := c.FormValue("text")
	expire := c.FormValue("expire")

	expire = converter.FormatingTime(expire)

	hashString := hash["hash"].(string)
	UploadText(hashString, text)

	database.DB.Create(&model.Text{TextHash: hashString, FileName: hashString + ".txt", ExpireAt: expire})

	return c.Render("templates/paste.html", hash)
}
