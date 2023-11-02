package handlers_articles

import (
	"github.com/gofiber/fiber/v2"
	"github.com/test_api_mm/database"
)

func CreateArticle_Handler(c *fiber.Ctx) error {
	var articles database.Article
	if err := c.BodyParser(&articles); err != nil {
		return fiber.ErrInternalServerError
	}

	if err := DB.Create(&articles).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(articles)
}
