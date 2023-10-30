package handlers_packages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/test_api_mm/database"
)

func CreatePackage_Handler(c *fiber.Ctx) error {
	var packages database.Package
	if err := c.BodyParser(&packages); err != nil {
		return fiber.ErrInternalServerError
	}

	if err := DB.Create(&packages).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(packages)
}
