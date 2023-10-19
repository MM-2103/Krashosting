package handlers_articles

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/test_api_mm/database"
)

var (
	DB *gorm.DB
)

func RetrieveAllArticles_Handler (c *fiber.Ctx) error {
	var articles []database.Article
	if err := DB.Find(&articles).Error; err != nil {
		fiber.ErrInternalServerError.Error()	
	}
    return c.JSON(articles)
}
