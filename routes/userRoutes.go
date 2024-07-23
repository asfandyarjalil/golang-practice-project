package routes

import (
	"github.com/asfandyarjalil/golang-practice-project/controllers"
	"github.com/asfandyarjalil/golang-practice-project/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, userController *controllers.UsersController) {
	app.Post("/login", middleware.ValidateCustomer, userController.Login)
}
