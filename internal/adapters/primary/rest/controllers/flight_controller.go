package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
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
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	flight, err := c.service.GetSpecificFlight(request)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Info().Msg("No flight found")
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Flight not found"})
		}
		log.Error().Err(err).Msg("Error getting specific flight")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting specific flight"})
	}

	return ctx.Status(http.StatusOK).JSON(flight)
}

func (c *FlightController) GetAllFlights(ctx *fiber.Ctx) error {
	flights, err := c.service.GetAllFlights()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all flights")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting all flights"})
	}

	if len(flights) == 0 {
		log.Info().Msg("No flights found")
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No flights available"})
	}

	return ctx.Status(http.StatusOK).JSON(flights)
}

func (c *FlightController) GetAllFlightsDestinationDateFlights(ctx *fiber.Ctx) error {
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
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No flights available"})
	}

	return ctx.Status(http.StatusOK).JSON(flights)
}

func (c *FlightController) GetAllActiveFlights(ctx *fiber.Ctx) error {
	flights, err := c.service.GetAllActiveFlights()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all active flights")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting all active flights"})
	}

	if len(flights) == 0 {
		log.Info().Msg("No active flights found")
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "There are no active flights"})
	}

	return ctx.Status(http.StatusOK).JSON(flights)
}

func (c *FlightController) CancelFlight(ctx *fiber.Ctx) error {
	var request entities.CancelFlightRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := c.service.CancelFlight(request)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Info().Msg("No flight found to cancel")
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No flight found to cancel"})
		}
		log.Error().Err(err).Msg("Error canceling flight")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error canceling flight"})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Flight canceled successfully"})
}

func (c *FlightController) AddFlight(ctx *fiber.Ctx) error {
	// Parse the request body
	var request entities.AddFlightRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Call the service to add the flight
	err := c.service.AddFlight(request)
	if err != nil {
		if err.Error() == "pq: insert or update on table \"flights\" violates foreign key constraint \"fk_plane_registration\"" {
			log.Error().Err(err).Msg("Foreign key constraint violation")
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid plane registration"})
		}
		log.Error().Err(err).Msg("Error adding flight")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error adding flight"})
	}

	log.Info().Msg("Successfully added flight")
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Flight added successfully"})
}

func (c *FlightController) FetchSeatMap(ctx *fiber.Ctx) error {
	var request entities.FetchSeatMapRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error parsing query parameters")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
	}

	seats, err := c.service.FetchSeatMap(request)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightID).Msg("Error fetching seat map")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching seat map"})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"seats": seats})
}
