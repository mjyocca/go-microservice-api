package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mjyocca/go-auth/backend/pkg/auth/routes"
)

func RegisterRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/login", routes.Login)
	auth.Post("/register", routes.Register)
}
