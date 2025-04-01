package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/core/services"
	"ams-service/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

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

func (c *FlightController) GetAllSpecificFlights(ctx *gin.Context) {
	// Extract user ID from the token
	userID, err := utils.ExtractIDFromToken(ctx, "user_id")
	if err != nil {
		log.Error().Err(err).Msg("Error extracting user ID from token")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Log the user ID for auditing purposes
	log.Info().Str("user_id", userID).Msg("User attempting to view specific flights")

	// Bind query parameters
	var request entities.GetSpecificFlightsRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query parameters")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	// Create channels for asynchronous processing
	resultChan := make(chan []entities.Flight)
	errorChan := make(chan error)

	// Run the service call in a Goroutine
	go func() {
		flights, err := c.service.GetAllSpecificFlights(request)
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
		log.Error().Err(err).Msg("Error getting specific flights")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting specific flights"})
	case <-time.After(10 * time.Second): // Timeout after 10 seconds
		log.Error().Msg("Request timed out")
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
	}
}

func (c *FlightController) GetAllActiveFlights(ctx *gin.Context) {
	// Extract employee ID from the token
	employeeID, err := utils.ExtractIDFromToken(ctx, "employee_id")
	if err != nil {
		log.Error().Err(err).Msg("Error extracting employee ID from token")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Log the employee ID for auditing purposes
	log.Info().Str("employee_id", employeeID).Msg("Employee attempting to view all active flights")

	// Create channels for asynchronous processing
	resultChan := make(chan []entities.Flight)
	errorChan := make(chan error)

	// Run the service call in a Goroutine
	go func() {
		flights, err := c.service.GetAllActiveFlights()
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
		log.Error().Err(err).Msg("Error getting all active flights")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting all active flights"})
	case <-time.After(10 * time.Second): // Timeout after 10 seconds
		log.Error().Msg("Request timed out")
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
	}
}

func (c *FlightController) CancelFlight(ctx *gin.Context) {
	// Extract employee ID from the token
	employeeID, err := utils.ExtractIDFromToken(ctx, "employee_id")
	if err != nil {
		log.Error().Err(err).Msg("Error extracting employee ID from token")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Log the employee ID for auditing purposes
	log.Info().Str("employee_id", employeeID).Msg("Employee attempting to cancel a flight")

	// Bind JSON request
	var request entities.CancelFlightRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Create channels for asynchronous processing
	resultChan := make(chan bool)
	errorChan := make(chan error)

	// Run the service call in a Goroutine
	go func() {
		err := c.service.CancelFlight(request)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- true
	}()

	// Wait for the result or error
	select {
	case <-resultChan:
		ctx.JSON(http.StatusOK, gin.H{"message": "Flight canceled successfully"})
	case err := <-errorChan:
		log.Error().Err(err).Msg("Error canceling flight")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error canceling flight"})
	case <-time.After(10 * time.Second): // Timeout after 10 seconds
		log.Error().Msg("Request timed out")
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
	}
}
