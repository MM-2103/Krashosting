package handlers_articles

import (
	"github.com/gofiber/fiber/v2"
	"github.com/test_api_mm/database"
)

func CreateArticle_Handler(c *fiber.Ctx) error {
	// Parse the request body to get the article data.
	var articles database.Article
	if err := c.BodyParser(&articles); err != nil {
		return fiber.ErrInternalServerError
	}

	// Insert the new article into the database.
	if err := DB.Create(&articles).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	// Return a success response with the created article.
	return c.Status(fiber.StatusCreated).JSON(articles)
}
