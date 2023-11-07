package auth_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/test_api_mm/database"
	"github.com/test_api_mm/validator"
)

type LoginBody struct {
	Username string `form:"username" validate:"required,max=255"`
	Password string `form:"password" validate:"required,min=8,max=32"`
}

func Login_Handler(c *fiber.Ctx) error {
	var body LoginBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	if err := validator.ValidateStruct(body); err != nil {
		return fiber.ErrBadRequest
	}
	var user database.User
	if err := database.DB.Where("username = ? OR username = ?", body.Username, body.Username).
		First(&user).Error; err != nil {
		return fiber.ErrUnauthorized
	}
	if valid := user.ValidatePassword(body.Password); !valid {
		return fiber.ErrUnauthorized
	}
	session_token := utils.UUID()
	access_token := utils.UUID()
	session := database.Session{
		UserID:       user.ID,
		SessionToken: session_token,
		AccessToken:  access_token,
		IP:           c.IP(),
	}
	if err := database.DB.Create(&session).Error; err != nil {
		return fiber.ErrInternalServerError
	}
	// Set cookies
	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    session_token,
		HTTPOnly: true,
	})
	// return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
	// 	"access_token": access_token,
	// })
	// Inside your signup handler, after the user is created successfully:
	return c.Redirect("/index.html") // Replace with the actual path to your homepage

}
