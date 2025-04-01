package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPassengerRoutes(router *gin.Engine, passengerController *controllers.PassengerController) {
	passengerRoute := router.Group("/passenger")
	{
		passengerRoute.POST("/checkin", passengerController.OnlineCheckInPassenger)
		passengerRoute.GET("/id", passengerController.GetPassengerByID)
		passengerRoute.GET("/pnr", passengerController.GetPassengerByPNR)
		passengerRoute.GET("/flight", passengerController.GetPassengersBySpecificFlight)
		passengerRoute.POST("/create", passengerController.CreatePassenger)
		passengerRoute.GET("/all", passengerController.GetAllPassengers)
	}
}
