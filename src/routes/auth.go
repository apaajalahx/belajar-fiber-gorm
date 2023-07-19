package routes

import (
	"dinarhamid/golanglearn/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App, controller *controllers.AuthHandler) {
	auth := app.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/register", controller.Register)
}
