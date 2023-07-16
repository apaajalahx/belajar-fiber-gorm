package main

import (
	"dinarhamid/golanglearn/src/controllers"
	"dinarhamid/golanglearn/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	UserHandler := controllers.UserController()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	routes.RegisterUser(app, UserHandler)

	app.Listen(":3000")
}
