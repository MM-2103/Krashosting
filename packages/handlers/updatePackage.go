package handlers_packages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/test_api_mm/database"
	"github.com/test_api_mm/validator"
	"gorm.io/gorm"
)

type UpdatePackage struct {
	Type          string
	Specefication string
}

func UpdatePackage_Handler(c *fiber.Ctx) error {
	packageID := c.Params("id")

	var updatedPackage UpdatePackage
	if err := c.BodyParser(&updatedPackage); err != nil {
		return fiber.ErrInternalServerError
	}
	if errors := validator.ValidateStruct(updatedPackage); errors != nil {
		return fiber.ErrBadRequest
	}

	var existingPackage database.Package
	if err := DB.First(&existingPackage, packageID).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return fiber.ErrInternalServerError
		}
	}
	existingPackage.Type = updatedPackage.Type
	existingPackage.Specefication = updatedPackage.Specefication
	if err := DB.Model(&existingPackage).Updates(existingPackage).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(existingPackage)
}
