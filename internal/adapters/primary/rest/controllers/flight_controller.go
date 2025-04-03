package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/utils"
	"net/http"

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

	log.Info().Str("employee_id", userID).Msg("User attempting to get a specific flight")

	flight, err := c.service.GetSpecificFlight(request, userID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Info().Msg("No flight found")
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "TODO:Flight not found"})
		}

		log.Error().Err(err).Msg("Error getting specific flight")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "TODO: Error getting specific flight"})
	}

	return ctx.Status(http.StatusOK).JSON(flight)
}

func (c *FlightController) GetAllFlights(ctx *fiber.Ctx) error {
	employeeID, err := utils.ExtractEmployeeIDFromToken(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error extracting employee ID from token")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	log.Info().Str("employee_id", employeeID).Msg("Employee attempting to view all flights")

	flights, err := c.service.GetAllFlights()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all flights")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting all flights"})
	}

	if len(flights) == 0 {
		log.Info().Msg("No flights found")
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "TODO: No flights available"})
	}

	return ctx.Status(http.StatusOK).JSON(flights)
}

func (c *FlightController) GetAllFlightsDestinationDateFlights(ctx *fiber.Ctx) error {
	userID, err := utils.ExtractIDFromToken(ctx, "user_id")
	if err != nil {
		log.Error().Err(err).Msg("Error extracting user ID from token")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	log.Info().Str("user_id", userID).Msg("User attempting to view specific flights")

	var request entities.GetAllFlightsDestinationDateRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query parameters")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
	}

	flights, err := c.service.GetAllFlightsDestinationDateFlights(request)
	if err != nil {
		log.Error().Err(err).Msg("Error getting specific flights")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting specific flights"})
	}

	if len(flights) == 0 {
		log.Info().Msg("No flights found")
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "TODO: No flights available"})
	}

	return ctx.Status(http.StatusOK).JSON(flights)
}

func (c *FlightController) GetAllActiveFlights(ctx *fiber.Ctx) error {
	employeeID, err := utils.ExtractIDFromToken(ctx, "employee_id")
	if err != nil {
		log.Error().Err(err).Msg("Error extracting employee ID from token")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	log.Info().Str("employee_id", employeeID).Msg("Employee attempting to view all active flights")

	flights, err := c.service.GetAllActiveFlights()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all active flights")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting all active flights"})
	}

	if len(flights) == 0 {
		log.Info().Msg("No active flights found")
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "TODO: There are no active flights"})
	}

	return ctx.Status(http.StatusOK).JSON(flights)
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

	err = c.service.CancelFlight(request)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Info().Msg("No flight found to cancel")
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "TODO: No flight found to cancel"})
		}
		log.Error().Err(err).Msg("Error canceling flight")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error canceling flight"})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Flight canceled successfully"})
}
