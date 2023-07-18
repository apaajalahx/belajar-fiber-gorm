package routes

import (
	"dinarhamid/golanglearn/src/controllers"
	"dinarhamid/golanglearn/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, controller *controllers.UserHandlers) {
	group := app.Group("users", middleware.Auth)
	group.Get("/", controller.GetUser)
}
