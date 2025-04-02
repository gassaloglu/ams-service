package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterFlightRoutes(app *fiber.App, flightController *controllers.FlightController) {
	flightRoute := app.Group("/flight")
	flightRoute.Get("/", flightController.GetSpecificFlight)
	flightRoute.Get("/all", flightController.GetAllFlights)
	flightRoute.Get("/all/", flightController.GetAllFlightsDestinationDateFlights)
	flightRoute.Get("/active", flightController.GetAllActiveFlights)
	flightRoute.Patch("/cancel", flightController.CancelFlight)
}
