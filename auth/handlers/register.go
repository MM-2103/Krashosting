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
	Firstname  string `form:"firstname" validate:"required,min=3,max=32"`
	Secondname string `form:"secondname" validate:"required,min=3,max=32"`
	Username   string `form:"username" validate:"required,min=3,max=32"`
	// Birthdate       string `form:"birthdate" validate:"datetime=mm-dd-yyyy"`
	Email           string `form:"email" validate:"required,email"`
	Password        string `form:"password" validate:"required,min=6,max=32"`
	Confirmpassword string `form:"confirmpassword" validate:"required,min=6,max=32"`
}

func Register_Handler(c *fiber.Ctx) error {
	var body RegisterBody
	if err := c.BodyParser(&body); err != nil {
		// Return the error message with details.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Can't parse the request body", "details": err.Error()})
	}

	validationErrors := validator.ValidateStruct(body)
	if validationErrors != nil {
		// Return detailed validation errors.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": validationErrors})
	}

	if DoesUserExist(body.Username, body.Email) {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		// Return the error message with details.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error encrypting password", "details": err.Error()})
	}

	user := database.User{
		Username: body.Username,
		Password: string(hash),
		Email:    body.Email,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		// Return the error message with details.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating user", "details": err.Error()})
	}

	// User created successfully.
	// return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
	// Inside your signup handler, after the user is created successfully:
	return c.Redirect("/index.html") // Replace with the actual path to your homepage

}
