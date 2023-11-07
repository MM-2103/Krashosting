package cart

import "github.com/gofiber/fiber/v2"

func Instance(app *fiber.App) {
	app.Post("/cart", func(c *fiber.Ctx) error {
		return nil
	})
}
