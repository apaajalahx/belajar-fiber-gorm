package controllers

import (
	"dinarhamid/golanglearn/src/models"
	"dinarhamid/golanglearn/system"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct{}

type UserCreateValidation struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserController() *UserHandlers {
	return &UserHandlers{}
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

	password, err := system.HashPassword(request.Password)
	if err != nil {
		fmt.Println(err)
	}

	model := models.UserModel{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
	}

	db := system.DatabaseHandler()
	result, err := db.CreateOne(model)
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(200).JSON(system.ResponseHandler("success", 200, result))

}
