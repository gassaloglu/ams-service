package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPlaneRoutes(app *fiber.App, planeController *controllers.PlaneController) {
	planeRoute := app.Group("/plane")
	planeRoute.Get("/all", planeController.GetAllPlanes)
	planeRoute.Post("/", planeController.AddPlane)
	planeRoute.Put("/status", planeController.SetPlaneStatus)
	planeRoute.Get("/:query", planeController.GetPlaneByRegistration)
	planeRoute.Get("/flightnumber", planeController.GetPlaneByFlightNumber)
	planeRoute.Get("/location", planeController.GetPlaneByLocation)
}
