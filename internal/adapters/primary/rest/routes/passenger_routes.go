package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPassengerRoutes(app *fiber.App, passengerController *controllers.PassengerController) {
	passengerRoute := app.Group("/passenger")
	passengerRoute.Get("/id", passengerController.GetPassengerByID)
	passengerRoute.Get("/pnr", passengerController.GetPassengerByPNR)
	passengerRoute.Post("/checkin", passengerController.OnlineCheckInPassenger)
	passengerRoute.Get("/specific-flight", passengerController.GetPassengersBySpecificFlight)
	passengerRoute.Post("/create", passengerController.CreatePassenger)
	passengerRoute.Get("/all", passengerController.GetAllPassengers)
	passengerRoute.Post("/employee-checkin", passengerController.EmployeeCheckInPassenger)
	passengerRoute.Patch("/cancel", passengerController.CancelPassenger)

}
