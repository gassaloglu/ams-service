package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterFlightRoutes(app *fiber.App, flightController *controllers.FlightController) {
	flightRoute := app.Group("/flight")
	flightRoute.Get("/specific", flightController.GetSpecificFlight)
	flightRoute.Get("/all", flightController.GetAllFlights)
	flightRoute.Get("/all-specific", flightController.GetAllSpecificFlights)
	flightRoute.Get("/all-active", flightController.GetAllActiveFlights)
	flightRoute.Post("/cancel", flightController.CancelFlight)
}
