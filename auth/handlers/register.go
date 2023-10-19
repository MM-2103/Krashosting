package auth_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/test_api_mm/database"
	"github.com/test_api_mm/validator"
	"golang.org/x/crypto/bcrypt"
)

func DoesUserExist(username string, email string) bool {
	if err := database.DB.Where("username = ? OR email = ?", username, email).First(&database.User{}).Error; err != nil {
		return false
	}
	return true
}

type RegisterBody struct {
	Username    string `json:"username" validate:"required,min=3,max=32"`
	PhoneNumber int    `json:"phone_number" validate:"required,min=10,max=10"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6,max=32"`
}

func Register_Handler(c *fiber.Ctx) error {
	var body RegisterBody
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}
	if errors := validator.ValidateStruct(body); errors != nil {
		return fiber.ErrBadRequest
	}
	if DoesUserExist(body.Username, body.Email) {
		return fiber.ErrConflict
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	user := database.User{
		Username:    body.Username,
		Password:    string(hash),
		Email:       body.Email,
		PhoneNumber: uint(body.PhoneNumber),
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}
