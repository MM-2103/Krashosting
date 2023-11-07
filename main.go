package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/test_api_mm/articles"
	"github.com/test_api_mm/auth"
	"github.com/test_api_mm/database"
)

func main() {
	database.Init()
	app := fiber.New()
	if !fiber.IsChild() {
		database.Migrate()
	}
	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.Dir("./frontend"),
		Browse: true,
	}))

	app.Use("/admin", filesystem.New(filesystem.Config{
		Root:   http.Dir("./admin"),
		Browse: true,
	}))

	// Default route to show when accessing the root
	auth.Instance(app)
	articles.Instance(app)

	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
