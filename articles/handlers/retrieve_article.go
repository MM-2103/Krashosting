package handlers_articles

import "github.com/gofiber/fiber/v2"

func RetrieveArticles_Handler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrInternalServerError
	}
	// Return actual response instead of nil
	return c.JSON(fiber.Map{
		"id": id,
	})

}
