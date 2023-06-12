package routes

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	// implement
	return c.SendString("Hello, World!")
}
