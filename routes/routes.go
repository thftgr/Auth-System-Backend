package routes

import (
	"Auth-System-Backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Get("/api/login", controllers.Login)
}
