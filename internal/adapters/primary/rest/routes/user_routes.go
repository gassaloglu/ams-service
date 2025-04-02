package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"
	"ams-service/internal/adapters/primary/rest/middlewares"
	"ams-service/internal/config"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, userController *controllers.UserController) {
	userRoute := app.Group("/user")
	{
		userRoute.Post("/register", userController.RegisterUser)
		userRoute.Post("/login", userController.LoginUser)
		userRoute.Use(middlewares.AuthMiddleware(config.JWTSecretKey))
	}
}
