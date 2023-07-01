package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Auth Middleware function to verify authorization exists
// checks token is still valid prior to accessing any protected routes
func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		/* check request is valid */
		if authorization == "" {
			return unauthorized(c)
		}

		token := strings.Split(authorization, "Bearer ")
		if len(token) < 2 {
			return unauthorized(c)
		}

		// to-do: validate token

		return unauthorized(c)
	}
}

func unauthorized(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusUnauthorized)
}
