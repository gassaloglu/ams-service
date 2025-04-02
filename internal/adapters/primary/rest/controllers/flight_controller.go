package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type FlightController struct {
	service primary.FlightService
}

func NewFlightController(service primary.FlightService) *FlightController {
	return &FlightController{service: service}
}

func (c *FlightController) GetSpecificFlight(ctx *fiber.Ctx) error {
	var request entities.GetSpecificFlightRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
	}

	userID, err := utils.ExtractUserIDFromToken(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error extracting user ID from token")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	resultChan := make(chan entities.Flight)
	errorChan := make(chan error)

	c.service.GetSpecificFlight(request, userID)

	select {
	case flight := <-resultChan:
		return ctx.Status(http.StatusOK).JSON(flight)
	case err := <-errorChan:
		log.Error().Err(err).Msg("Error getting specific flight")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting specific flight"})
	case <-time.After(10 * time.Second):
		log.Error().Msg("Request timed out")
		return ctx.Status(http.StatusGatewayTimeout).JSON(fiber.Map{"error": "Request timed out"})
	}
}

func (c *FlightController) GetAllFlights(ctx *fiber.Ctx) error {
	employeeID, err := utils.ExtractEmployeeIDFromToken(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error extracting employee ID from token")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	log.Info().Str("employee_id", employeeID).Msg("Employee attempting to view all flights")

	resultChan := make(chan []entities.Flight)
	errorChan := make(chan error)

	go func() {
		flights, err := c.service.GetAllFlights()
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- flights
	}()

	select {
	case flights := <-resultChan:
		return ctx.Status(http.StatusOK).JSON(flights)
	case err := <-errorChan:
		log.Error().Err(err).Msg("Error getting all flights")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting all flights"})
	case <-time.After(10 * time.Second):
		log.Error().Msg("Request timed out")
		return ctx.Status(http.StatusGatewayTimeout).JSON(fiber.Map{"error": "Request timed out"})
	}
}

func (c *FlightController) GetAllSpecificFlights(ctx *fiber.Ctx) error {
	userID, err := utils.ExtractIDFromToken(ctx, "user_id")
	if err != nil {
		log.Error().Err(err).Msg("Error extracting user ID from token")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	log.Info().Str("user_id", userID).Msg("User attempting to view specific flights")

	var request entities.GetSpecificFlightsRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query parameters")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
	}

	resultChan := make(chan []entities.Flight)
	errorChan := make(chan error)

	go func() {
		flights, err := c.service.GetAllSpecificFlights(request)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- flights
	}()

	select {
	case flights := <-resultChan:
		return ctx.Status(http.StatusOK).JSON(flights)
	case err := <-errorChan:
		log.Error().Err(err).Msg("Error getting specific flights")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting specific flights"})
	case <-time.After(10 * time.Second):
		log.Error().Msg("Request timed out")
		return ctx.Status(http.StatusGatewayTimeout).JSON(fiber.Map{"error": "Request timed out"})
	}
}

func (c *FlightController) GetAllActiveFlights(ctx *fiber.Ctx) error {
	employeeID, err := utils.ExtractIDFromToken(ctx, "employee_id")
	if err != nil {
		log.Error().Err(err).Msg("Error extracting employee ID from token")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	log.Info().Str("employee_id", employeeID).Msg("Employee attempting to view all active flights")

	resultChan := make(chan []entities.Flight)
	errorChan := make(chan error)

	go func() {
		flights, err := c.service.GetAllActiveFlights()
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- flights
	}()

	select {
	case flights := <-resultChan:
		return ctx.Status(http.StatusOK).JSON(flights)
	case err := <-errorChan:
		log.Error().Err(err).Msg("Error getting all active flights")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting all active flights"})
	case <-time.After(10 * time.Second):
		log.Error().Msg("Request timed out")
		return ctx.Status(http.StatusGatewayTimeout).JSON(fiber.Map{"error": "Request timed out"})
	}
}

func (c *FlightController) CancelFlight(ctx *fiber.Ctx) error {
	employeeID, err := utils.ExtractIDFromToken(ctx, "employee_id")
	if err != nil {
		log.Error().Err(err).Msg("Error extracting employee ID from token")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	log.Info().Str("employee_id", employeeID).Msg("Employee attempting to cancel a flight")

	var request entities.CancelFlightRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resultChan := make(chan bool)
	errorChan := make(chan error)

	go func() {
		err := c.service.CancelFlight(request)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- true
	}()

	select {
	case <-resultChan:
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Flight canceled successfully"})
	case err := <-errorChan:
		log.Error().Err(err).Msg("Error canceling flight")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error canceling flight"})
	case <-time.After(10 * time.Second):
		log.Error().Msg("Request timed out")
		return ctx.Status(http.StatusGatewayTimeout).JSON(fiber.Map{"error": "Request timed out"})
	}
}
