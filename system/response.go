package system

import (
	"github.com/gofiber/fiber/v2"
)

type Array interface {
	interface{} | []string | []int
}

func ResponseHandler[T Array](message string, code int16, data T) fiber.Map {
	return fiber.Map{
		"status_code": code,
		"message":     message,
		"data":        data,
	}
}
