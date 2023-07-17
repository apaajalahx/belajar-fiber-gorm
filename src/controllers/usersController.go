package controllers

import (
	"dinarhamid/golanglearn/system"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct{}

func UserController() *UserHandlers {
	return &UserHandlers{}
}

func (h *UserHandlers) GetUser(c *fiber.Ctx) error {
	fmt.Println(os.Getenv("port"))
	return c.JSON(system.ResponseHandler("success", 200, []string{}))
}
