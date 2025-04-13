package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"
	"ams-service/internal/adapters/primary/rest/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterEmployeeRoutes(app *fiber.App, employeeController *controllers.EmployeeController) {
	employeeRoute := app.Group("/employee")

	// Public routes
	employeeRoute.Post("/login", employeeController.LoginEmployee)
	employeeRoute.Post("/register", employeeController.RegisterEmployee)

	// Protected routes
	employeeRoute.Use(middlewares.ProtectionForEmployees())
	employeeRoute.Get("/", employeeController.GetEmployeeByID)
}
