package handlers_articles

import (
	"github.com/gofiber/fiber/v2"
	"github.com/test_api_mm/database"
)

func DeleteArticle_Handler(c *fiber.Ctx) error {
	articleID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var existingArticle database.Article
	if err := DB.First(&existingArticle, articleID).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	if err := DB.Delete(&existingArticle).Error; err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
