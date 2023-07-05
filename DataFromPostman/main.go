package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RequestData struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func main() {
	app := fiber.New()
	app.Post("/api", func(c *fiber.Ctx) error {
		data := new(RequestData)

		if err := c.BodyParser(data); err != nil {
			return err
		}

		// Access the bound data
		fmt.Println(data.Name)
		fmt.Println(data.Email)
		fmt.Println(data.Message)

		return nil
	})

	app.Listen(":3000") // Change the port as needed

}
