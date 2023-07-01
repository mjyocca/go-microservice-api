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
	// var user models.User

	var reqUser Credentials
	if err := c.BodyParser(&reqUser); err != nil {
		return err
	}

	fmt.Printf("\n%+v\n", reqUser)

	return c.SendString("Hello, World!")
}
