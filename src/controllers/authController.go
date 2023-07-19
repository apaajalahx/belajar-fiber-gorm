package controllers

import (
	"dinarhamid/golanglearn/src/models"
	"dinarhamid/golanglearn/system"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler struct {
	model *models.UserModelService
}

func AuthController(db *gorm.DB) *AuthHandler {
	model := models.UserModelHandler(db)
	return &AuthHandler{
		model: model,
	}
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

type Register struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var request Register

	c.BodyParser(&request)

	res, errs := h.model.GetOne(&models.UserModel{
		Email: request.Email,
	})
	fmt.Println(res)
	if errors.Is(errs, gorm.ErrRecordNotFound) {

		pwd, _ := system.HashPassword(request.Password)

		res, err := h.model.CreateOne(&models.UserModel{
			Email:    request.Email,
			Username: request.Username,
			Name:     request.Name,
			Password: pwd,
		})

		if err != nil {
			c.Status(500).JSON(system.ResponseHandler(err.Error(), 500, []string{}))
		}

		return c.Status(200).JSON(system.ResponseHandler("success", 200, res))
	}

	return c.Status(500).JSON(system.ResponseHandler("users with same email has been registered", 500, []string{}))
}
