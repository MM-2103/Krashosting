package auth

import (
	"github.com/gofiber/fiber/v2"
	auth_handlers "github.com/test_api_mm/auth/handlers"
)

func Instance(app *fiber.App) {
	auth := app.Group("/auth")
	{
		auth.Post("/login", auth_handlers.Login_Handler)
	}
}
