package auth

import "github.com/gofiber/fiber/v2"

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return unauthorized(c)
	}
}

func unauthorized(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusUnauthorized)
}
