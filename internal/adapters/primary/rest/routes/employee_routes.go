package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"
	"ams-service/internal/adapters/primary/rest/middlewares"
	"ams-service/internal/config"

	"github.com/gofiber/fiber/v2"
)

func RegisterEmployeeRoutes(app *fiber.App, employeeController *controllers.EmployeeController) {
	employeeRoute := app.Group("/employee")

	// Public routes
	employeeRoute.Post("/login", employeeController.LoginEmployee)
	employeeRoute.Post("/register", employeeController.RegisterEmployee)

	// Protected routes (require AuthMiddleware)
	employeeRoute.Use(middlewares.AuthMiddleware(config.JWTSecretKey))
	employeeRoute.Get("/", employeeController.GetEmployeeByID)
}
