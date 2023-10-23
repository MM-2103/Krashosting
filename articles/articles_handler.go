package articles

import (
	"github.com/gofiber/fiber/v2"
	handlers_articles "github.com/test_api_mm/articles/handlers"
)

func Instance(app *fiber.App) {
	articles := app.Group("/articles")
	{
		articles.Get("/:id<int>", handlers_articles.RetrieveArticles_Handler)
		articles.Get("/", handlers_articles.RetrieveArticles_Handler)
		articles.Post("", handlers_articles.CreateArticle_Handler)
		articles.Delete("/:id<int>", handlers_articles.DeleteArticle_Handler)
		articles.Put("/:id<int>", handlers_articles.UpdateArticle_Handler)
	}
}
