package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type RequestData struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func main() {

	dsn := "host=localhost user=postgres password=athun123 dbname=sigleapitest port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Db = db
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&RequestData{})
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Post("test", test)
	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}

func test(c *fiber.Ctx) error {
	data := new(RequestData)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	resutl := Db.Create(&data)
	if resutl.Error != nil {
		log.Fatal(resutl.Error)
	}

	return nil
}
