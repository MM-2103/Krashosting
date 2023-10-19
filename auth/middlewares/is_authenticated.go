package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

var (
	AllowedRoutes = []string{
		"/auth/login",
		"/auth/register",
	}
)

func IsAuthenticated_Middleware(c *fiber.Ctx) error {
	cookie := c.Cookies("session_token")
	if cookie == "" {
		if slices.Contains() {
		}
		return fiber.ErrUnauthorized
	}
}
