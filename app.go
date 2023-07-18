package main

import (
	"dinarhamid/golanglearn/src/controllers"
	"dinarhamid/golanglearn/src/routes"
	"dinarhamid/golanglearn/system"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	err := system.LoadConfig("config.json")

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	UserHandler := controllers.UserController()
	AuthHandler := controllers.AuthController()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(os.Getenv("port"))
	})

	routes.RegisterUserRoutes(app, UserHandler)
	routes.RegisterAuthRoutes(app, AuthHandler)

	app.Listen(":3000")
}
