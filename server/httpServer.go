package server

import (
	"database/sql"
	"log"

	"github.com/asfandyarjalil/golang-practice-project/controllers"
	"github.com/asfandyarjalil/golang-practice-project/repositories"
	"github.com/asfandyarjalil/golang-practice-project/routes"
	"github.com/asfandyarjalil/golang-practice-project/services"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config *viper.Viper
	app    *fiber.App
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	userRepository := repositories.NewUsersRepository(dbHandler)
	userService := services.NewUsersService(userRepository)
	userController := controllers.NewUsersController(userService)

	app := fiber.New()

	// routes
	routes.UserRoute(app, userController)

	return HttpServer{config: config, app: app}
}

func (hs HttpServer) Start() {
	err := hs.app.Listen(hs.config.GetString("SERVER_PORT"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
