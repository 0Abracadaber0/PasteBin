package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
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
	//expire := c.FormValue("expire")

	hashString := hash["hash"].(string)

	UploadText(hashString, text)

	return c.Render("templates/paste.html", hash)
}
