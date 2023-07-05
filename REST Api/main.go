package main

import (
	"log"

	"github.com/athunlal/api/config"
	db "github.com/athunlal/api/datatabase"
	handler "github.com/athunlal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db.InitDB(c.DBUrl)
	app := fiber.New()
	app.Post("register", handler.Register)
	app.Listen(":3000")
}
