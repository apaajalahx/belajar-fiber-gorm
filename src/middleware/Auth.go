package middleware

import (
	"dinarhamid/golanglearn/system"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	bearer := c.Get("Authorization")
	token := strings.Split(bearer, " ")
	if len(token) > 1 {
		res, err := system.Verify(token[1])
		if err != nil {
			return c.Status(401).JSON(system.ResponseHandler("Unauthorize", 401, []string{}))
		}

		c.Locals("users", res)

		return c.Next()
	}
	return c.Status(401).JSON(system.ResponseHandler("Bearer not found", 401, []string{}))
}
