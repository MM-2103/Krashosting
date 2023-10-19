package articles

import (
	"github.com/gofiber/fiber/v2"
	handlers_articles "github.com/test_api_mm/articles/handlers"
)

func Instance(app *fiber.App) {
	articles := app.Group("/articles")
	{
		articles.Get("/:id<int>", handlers_articles.RetrieveArticles_Handler)
	}
}
