package controllers

import (
	"ams-service/core/entities"
	"ams-service/core/services"
	"ams-service/middlewares"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var FLIGHT_LOG_PREFIX string = "flight_controller.go"

type FlightController struct {
	service *services.FlightService
}

func NewFlightController(service *services.FlightService) *FlightController {
	return &FlightController{service: service}
}

func (c *FlightController) GetSpecificFlight(ctx *gin.Context) {
	var request entities.GetSpecificFlightRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error binding query: %v", FLIGHT_LOG_PREFIX, err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "TODO: Error binding query"})
		return
	}

	flight, err := c.service.GetSpecificFlight(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting specific flight: %v", FLIGHT_LOG_PREFIX, err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "TODO: Error getting specific flight"})
		return
	}
	ctx.JSON(http.StatusOK, flight)
}

func (c *FlightController) GetAllFlights(ctx *gin.Context) {
	flights, err := c.service.GetAllFlights()
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting all flights: %v", FLIGHT_LOG_PREFIX, err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "TODO: Error getting all flights"})
		return
	}
	ctx.JSON(http.StatusOK, flights)
}
