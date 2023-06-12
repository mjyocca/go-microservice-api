package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func Login(c *fiber.Ctx) error {

	var creds Credentials

	if err := c.BodyParser(&creds); err != nil {
		return err
	}

	fmt.Printf("\n%+v\n", creds)

	return c.SendString("Hello, World!")
}
