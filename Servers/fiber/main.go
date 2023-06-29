package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define routes and their handlers
	app.Get("/", helloWorld)
	app.Get("/user", getUser)

	// Start the server
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func getUser(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"user": "This is JSON respose",
	})
}
