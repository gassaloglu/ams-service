package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPlaneRoutes(app *fiber.App, planeController *controllers.PlaneController) {
	planeRoute := app.Group("/planes")
	planeRoute.Get("/", planeController.GetAllPlanes)
	planeRoute.Post("/", planeController.CreatePlane)
	planeRoute.Patch("/:registration", planeController.SetPlaneStatus)
	planeRoute.Get("/:registration", planeController.GetPlaneByRegistration)
}
