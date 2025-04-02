package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPassengerRoutes(app *fiber.App, passengerController *controllers.PassengerController) {
	passengerRoute := app.Group("/passenger")
	passengerRoute.Get("/id", passengerController.GetPassengerByID)
	passengerRoute.Get("/pnr", passengerController.GetPassengerByPNR)
	passengerRoute.Post("/check-in", passengerController.OnlineCheckInPassenger)
	passengerRoute.Get("/flight", passengerController.GetPassengersBySpecificFlight)
	passengerRoute.Post("/create", passengerController.CreatePassenger)
	passengerRoute.Get("/all", passengerController.GetAllPassengers)
}
