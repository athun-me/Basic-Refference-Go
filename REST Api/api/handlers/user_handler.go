package handler

import (
	"fmt"
	"log"

	database "github.com/athunlal/api/datatabase"
	"github.com/athunlal/api/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	Data := new(models.User)
	if err := c.BodyParser(Data); err != nil {
		fmt.Println("Body parsing error")
		return err
	}

	result := database.DB.Create(&Data)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return nil
}
