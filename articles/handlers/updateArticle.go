package handlers_articles

import (
	"github.com/gofiber/fiber/v2"
	"github.com/test_api_mm/database"
	"github.com/test_api_mm/validator"
	"gorm.io/gorm"
)

type UpdateBody struct {
	Title string
	Body  string
}

func UpdateArticle_Handler(c *fiber.Ctx) error {
	articleID := c.Params("id")

	var updatedArticle UpdateBody
	if err := c.BodyParser(&updatedArticle); err != nil {
		return fiber.ErrInternalServerError
	}
	if errors := validator.ValidateStruct(updatedArticle); errors != nil {
		return fiber.ErrBadRequest
	}

	var existingArticle database.Article
	if err := DB.First(&existingArticle, articleID).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return fiber.ErrInternalServerError
		}
	}
	existingArticle.Title = updatedArticle.Title
	existingArticle.Body = updatedArticle.Body
	if err := DB.Model(&existingArticle).Updates(existingArticle).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(existingArticle)
}
