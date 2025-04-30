package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, userController *controllers.UserController) {
	userRoute := app.Group("/users")
	userRoute.Post("/", userController.RegisterUser)
	userRoute.Post("/sessions", userController.LoginUser)
}
