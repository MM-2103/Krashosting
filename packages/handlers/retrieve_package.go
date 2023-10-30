package handlers_packages

import "github.com/gofiber/fiber/v2"

func RetrievePackages_Handler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}
