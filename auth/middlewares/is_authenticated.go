package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/test_api_mm/database"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
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
		if slices.Contains(AllowedRoutes, c.Path()) {
			return c.Next()
		}
		return fiber.ErrUnauthorized
	}
	var session database.Session
	if err := database.DB.Where("session_token = ? AND valid = ?", cookie, true).First(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if slices.Contains(AllowedRoutes, c.Path()) {
				return c.Next()
			}
			return fiber.ErrUnauthorized
		}
		return fiber.ErrInternalServerError
	}
	c.Locals("session", session)
	return c.Next()
}
