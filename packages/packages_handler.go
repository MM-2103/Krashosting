package packages

import (
	"github.com/gofiber/fiber/v2"
)

func Instance(app *fiber.App) {
	packages := app.Group("/packages")
	{
		packages.Get("/:id<int>", handlers_packages)
	}
}
