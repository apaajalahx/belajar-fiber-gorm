package main

import (
	"dinarhamid/golanglearn/src/controllers"
	"dinarhamid/golanglearn/src/models"
	"dinarhamid/golanglearn/src/routes"
	"dinarhamid/golanglearn/system"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	err := system.LoadConfig("config.json")

	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// migration database
	db.AutoMigrate(&models.UserModel{})

	app := fiber.New()

	// register Routes
	UserHandler := controllers.UserController(db)
	AuthHandler := controllers.AuthController(db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(os.Getenv("port"))
	})

	routes.RegisterUserRoutes(app, UserHandler)
	routes.RegisterAuthRoutes(app, AuthHandler)

	app.Listen(":3000")
}
