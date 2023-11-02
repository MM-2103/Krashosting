package handlers_packages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/test_api_mm/database"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func RetrieveAllPackages_Handler(c *fiber.Ctx) error {
	var packages []database.Package
	if err := DB.Find(&packages).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve articles",
		})
	}
	return c.JSON(packages)
}
