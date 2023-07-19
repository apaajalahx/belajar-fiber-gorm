package controllers

import (
	"dinarhamid/golanglearn/src/models"
	"dinarhamid/golanglearn/system"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandlers struct {
	model *models.UserModelService
}

type UserCreateValidation struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserController(db *gorm.DB) *UserHandlers {
	model := models.UserModelHandler(db)
	return &UserHandlers{
		model: model,
	}
}

func (h *UserHandlers) GetUser(c *fiber.Ctx) error {
	return c.JSON(system.ResponseHandler("success", 200, []string{}))
}

func (h *UserHandlers) CreateOne(c *fiber.Ctx) error {

	var request UserCreateValidation

	err := c.BodyParser(&request)

	if err != nil {
		return c.Status(422).JSON(system.ResponseHandler("Error validation"+err.Error(), 422, []string{}))
	}

	pwd, err := system.HashPassword(request.Password)
	if err != nil {
		fmt.Println(err)
	}

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
