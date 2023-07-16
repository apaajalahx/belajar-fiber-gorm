package routes

import (
	"dinarhamid/golanglearn/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(app *fiber.App, controller *controllers.UserHandlers) {
	group := app.Group("users")
	group.Get("/", controller.GetUser)
}
