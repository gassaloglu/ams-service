package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"
	"ams-service/internal/adapters/primary/rest/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterFlightRoutes(app *fiber.App, flightController *controllers.FlightController) {
	flightRoute := app.Group("/flights")

	// Generally available routes for authenticated requests
	{
		group := flightRoute.Group("/")
		group.Use(middlewares.Protection())

		group.Get("/", flightController.GetAllActiveFlights)
		group.Get("/:id", flightController.GetFlightById)
	}

	// Admin, flight planner, passenger services, and ground services routes
	{
		group := flightRoute.Group("/")

		group.Use(middlewares.ProtectionForRoles([]string{
			middlewares.AdminRole,
			middlewares.FlightPlannerRole,
			middlewares.PassengerServicesRole,
			middlewares.GroundServicesRole},
		))

		// group.Get("/", flightController.GetAllFlights)
	}

	// Admin and flight planner routes
	{
		group := flightRoute.Group("/")

		group.Use(middlewares.ProtectionForRoles([]string{
			middlewares.AdminRole,
			middlewares.FlightPlannerRole},
		))

		group.Post("/", flightController.CreateFlight)
	}
}
