package handlers_articles

import (
	"github.com/gofiber/fiber/v2"
	"github.com/test_api_mm/database"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func RetrieveAllArticles_Handler(c *fiber.Ctx) error {
	var articles []database.Article
	if err := DB.Find(&articles).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve articles",
		})
	}
	return c.JSON(articles)
}
