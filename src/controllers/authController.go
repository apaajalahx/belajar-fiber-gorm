package controllers

import (
	"dinarhamid/golanglearn/src/models"
	"dinarhamid/golanglearn/system"
	"errors"

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

	var login Login

	c.BodyParser(&login)

	result, err := h.model.GetOne(&models.UserModel{
		Username: login.Username,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(404).JSON(system.ResponseHandler(err.Error(), 404, []string{}))
		}
	}

	if !system.CheckPassword(login.Password, result.Password) {
		c.Status(401).JSON(system.ResponseHandler("Wrong Username or password", 401, []string{}))
	}

	token, err := system.Sign(result.Id, result.Username, result.Password)
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

	_, errs := h.model.GetOne(&models.UserModel{
		Email: request.Email,
	})

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
