package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mjyocca/go-auth/backend/pkg/auth"
)

func newApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	return app
}

func main() {

	app := newApp()

	app.Get("/health")

	auth.RegisterService(app)

	if err := app.Listen(":8000"); err != nil {
		fmt.Printf("%v", err)
	}
}
