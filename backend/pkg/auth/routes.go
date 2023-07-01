package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mjyocca/go-auth/backend/pkg/auth/db"
	"github.com/mjyocca/go-auth/backend/pkg/auth/routes"
	service "github.com/mjyocca/go-auth/backend/pkg/auth/services"
)

// rename to RegisterService
func RegisterService(app *fiber.App) {

	// to-do: add db url
	handler := db.NewHandler("")
	// to-do: return service to inject into routes/controller
	service.NewService(handler)

	// to-do: pass service to routes/controller
	// r := routes.NewRouter(service)
	// auth.Post("/login", r.Login)
	auth := app.Group("/auth")
	auth.Post("/login", routes.Login)
	auth.Post("/register", routes.Register)
}
