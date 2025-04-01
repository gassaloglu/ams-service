package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterFlightRoutes(router *gin.Engine, flightController *controllers.FlightController) {
	flightRoute := router.Group("/flight")
	{
		flightRoute.GET("/", flightController.GetSpecificFlight)
		flightRoute.GET("/all", flightController.GetAllFlights)
		flightRoute.GET("/all/", flightController.GetAllSpecificFlights)
		flightRoute.GET("/active", flightController.GetAllActiveFlights)
		flightRoute.PATCH("/cancel", flightController.CancelFlight)
	}
}
