package handlers_packages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/test_api_mm/database"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DeletePackage_Handler(c *fiber.Ctx) error {
	packageID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var existingPackage database.Package
	if err := DB.First(&existingPackage, packageID).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	if err := DB.Delete(&existingPackage).Error; err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
