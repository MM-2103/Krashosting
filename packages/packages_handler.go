package packages

import (
	"github.com/gofiber/fiber/v2"
	handlers_packages "github.com/test_api_mm/packages/handlers"
)

func Instance(app *fiber.App) {
	packages := app.Group("/packages")
	{
		packages.Get("/:id<int>", handlers_packages.RetrieveAllPackages_Handler)
		packages.Get("/", handlers_packages.RetrievePackages_Handler)
		packages.Post("", handlers_packages.CreatePackage_Handler)
		packages.Delete("/:id<int>", handlers_packages.DeletePackage_Handler)
		packages.Put("/:id<int>", handlers_packages.UpdatePackage_Handler)
	}
}
