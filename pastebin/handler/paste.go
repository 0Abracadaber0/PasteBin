package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Ticket struct {
	Number int `json:"number"`
}

func PasteHandler(c *fiber.Ctx) error {
	url := "http://hashgenerator:8081/api/calc_hash"

	ticket := Ticket{
		Number: 123,
	}

	jsonData, err := json.Marshal(ticket)
	if err != nil {
		log.Println(err)
	}

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)
	req.Header.SetContentType("application/json")
	req.SetBody(jsonData)

	if err = a.Parse(); err != nil {
		panic(err)
	}

	_, body, errs := a.Bytes()
	if errs != nil {
		log.Fatal(errs)
	}

	var data fiber.Map
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return c.Render("templates/paste.html", data)
}
