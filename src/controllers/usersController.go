package controllers

import (
	"dinarhamid/golanglearn/system"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct{}

func UserController() *UserHandlers {
	return &UserHandlers{}
}

func (h *UserHandlers) GetUser(c *fiber.Ctx) error {
	return c.JSON(system.ResponseHandler("success", 200, []string{}))
}
