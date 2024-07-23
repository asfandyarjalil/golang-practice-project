package middleware

import (
	"github.com/asfandyarjalil/golang-practice-project/models"
	"github.com/gofiber/fiber/v2"
)

// ValidateCustomer validates the customer struct
func ValidateCustomer(c *fiber.Ctx) error {
	var user models.LoginRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"success": false,
			"status":  fiber.StatusBadRequest,
		})
	}

	if user.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username is required",
			"success": false,
			"status":  fiber.StatusBadRequest,
		})

	}
	if user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password is required",
			"success": false,
			"status":  fiber.StatusBadRequest,
		})
	}
	// Continue to the next handler if validation passes
	return c.Next()
}
