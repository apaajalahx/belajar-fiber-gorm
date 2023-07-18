package controllers

import (
	"dinarhamid/golanglearn/system"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{}

func AuthController() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	token, err := system.Sign("DinarHamid123", "palkna3@gmail.com")
	if err != nil {
		return c.Status(500).JSON(system.ResponseHandler(err.Error(), 500, []string{}))
	}
	return c.Status(200).JSON(system.ResponseHandler("oke", 200, map[string]any{
		"token":     token,
		"expiresIn": 3600,
		"type":      "Bearer Token",
	}))
}
