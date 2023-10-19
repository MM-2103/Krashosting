package main

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/test_api_mm/articles"
	"github.com/test_api_mm/database"
)

var embedDirStatic embed.FS

func main() {
	database.Init()
	app := fiber.New()
	if !fiber.IsChild() {
		database.Migrate()
	}
	app.Use("/krashosting", filesystem.New(filesystem.Config{
		Root: http.FS(
			embedDirStatic,
		),
		Browse: true,
	}))

	// Default route to show when accessing the root
	articles.Instance(app)

	app.Listen(":3000")
}
