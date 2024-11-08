package main

import (
	"ams-service/core/services"
	"ams-service/infrastructure/api/controllers"
	"ams-service/infrastructure/persistence/repositories"
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
