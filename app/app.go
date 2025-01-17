package app

import (
	"ams-service/core/services"
	"ams-service/infrastructure/api/controllers"
	"ams-service/infrastructure/persistence/repositories"
	"ams-service/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

var LOG_PREFIX string = "app.go"

func Run() {
	// Initialize repositories
	passengerRepo := repositories.NewPassengerRepositoryImpl()

	// Initialize services
	passengerService := services.NewPassengerService(passengerRepo)

	// Initialize controllers
	passengerController := controllers.NewPassengerController(passengerService)

	// Setup router
	router := gin.Default()
	router.Use(middlewares.Logger())
	router.Use(middlewares.ErrorHandler())

	// Setup routes
	passengerRoute := router.Group("/passenger")
	{
		passengerRoute.POST("/checkin", passengerController.OnlineCheckInPassenger)
		passengerRoute.GET("/:id", passengerController.GetPassengerByID)
	}

	// Run the server
	err := router.Run(":8080")
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Failed to start server: %v", LOG_PREFIX, err))

	}
}
