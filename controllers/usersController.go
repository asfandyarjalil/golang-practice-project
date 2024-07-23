package controllers

import (
	"github.com/asfandyarjalil/golang-practice-project/models"
	"github.com/asfandyarjalil/golang-practice-project/services"
	"github.com/gofiber/fiber/v2"
)

type UsersController struct {
	userService *services.UsersService
}

func NewUsersController(userService *services.UsersService) *UsersController {
	return &UsersController{
		userService: userService,
	}
}

func (uc *UsersController) Login(ctx *fiber.Ctx) error {
	bodyData := new(models.LoginRequest)

	if err := ctx.BodyParser(bodyData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}
	accessToken, responseError := uc.userService.Login(bodyData.Username, bodyData.Password)
	if responseError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseError)
	}
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":     "Login Successfully",
		"accessToken": accessToken,
	})
}
