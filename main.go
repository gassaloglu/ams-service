package main

import (
	"ams-passenger-service/core/services"
	"ams-passenger-service/infrastructure/api/controllers"
	"ams-passenger-service/infrastructure/persistence/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	passengerRepo := repositories.NewPassengerRepositoryImpl()

	passengerService := services.NewPassengerService(passengerRepo)

	passengerController := controllers.NewPassengerController(passengerService)

	router := gin.Default()
	passengerRoute := router.Group("/passenger")
	{
		passengerRoute.POST("/checkin", passengerController.OnlineCheckInPassenger)
	}

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
