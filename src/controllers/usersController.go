package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct{}

func UserController() *UserHandlers {
	return &UserHandlers{}
}

func (h *UserHandlers) GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"name":        "Dinar Hamid",
		"username":    "dinar1337",
		"email":       "palkna3@gmail.com",
		"is_verified": true,
	})
}
