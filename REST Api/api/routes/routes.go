package routes

import (
	handler "github.com/athunlal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {

	app.Post("register", handler.Register)
}
