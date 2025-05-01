package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"
	"ams-service/internal/adapters/primary/rest/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterEmployeeRoutes(app *fiber.App, employeeController *controllers.EmployeeController) {
	employeeRoute := app.Group("/employees")

	// Public routes
	employeeRoute.Post("/sessions", employeeController.LoginEmployee)

	// Protected routes
	employeeRoute.Use(middlewares.ProtectionForEmployees())
	employeeRoute.Get("/", employeeController.GetEmployees)

	// Admin routes
	employeeRoute.Use(middlewares.ProtectionForRoles([]string{"admin"}))
	employeeRoute.Post("/", employeeController.RegisterEmployee)
}
