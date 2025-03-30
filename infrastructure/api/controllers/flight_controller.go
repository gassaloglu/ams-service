package controllers

import (
	"ams-service/core/entities"
	"ams-service/core/services"
	"ams-service/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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
		log.Error().Err(err).Msg("Error binding query")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	userID, err := utils.ExtractUserIDFromToken(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error extracting user ID from token")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	resultChan := make(chan entities.Flight)
	errorChan := make(chan error)

	c.service.GetSpecificFlight(request, userID, resultChan, errorChan)

	select {
	case flight := <-resultChan:
		ctx.JSON(http.StatusOK, flight)
	case err := <-errorChan:
		log.Error().Err(err).Msg("Error getting specific flight")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting specific flight"})
	case <-time.After(10 * time.Second):
		log.Error().Msg("Request timed out")
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
	}
}

func (c *FlightController) GetAllFlights(ctx *gin.Context) {
	// Extract employee ID from the token
	employeeID, err := utils.ExtractEmployeeIDFromToken(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error extracting employee ID from token")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Log the employee ID for auditing purposes
	log.Info().Str("employee_id", employeeID).Msg("Employee attempting to view all flights")

	// Create channels for asynchronous processing
	resultChan := make(chan []entities.Flight)
	errorChan := make(chan error)

	// Run the service call in a Goroutine
	go func() {
		flights, err := c.service.GetAllFlights()
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- flights
	}()

	// Wait for the result or error
	select {
	case flights := <-resultChan:
		ctx.JSON(http.StatusOK, flights)
	case err := <-errorChan:
		log.Error().Err(err).Msg("Error getting all flights")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting all flights"})
	case <-time.After(10 * time.Second):
		log.Error().Msg("Request timed out")
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
	}
}
