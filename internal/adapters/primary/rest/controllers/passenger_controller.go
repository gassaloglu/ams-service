package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type PassengerController struct {
	service primary.PassengerService
}

func NewPassengerController(service primary.PassengerService) *PassengerController {
	return &PassengerController{service: service}
}

func (c *PassengerController) GetPassengerByID(ctx *fiber.Ctx) error {
	var request entities.GetPassengerByIdRequest
	if err := ctx.QueryParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	passenger, err := c.service.GetPassengerByID(request)
	if err != nil {
		log.Error().Err(err).Str("national_id", request.NationalId).Msg("Error getting passenger by ID")
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Passenger not found"})
	}
	log.Info().Str("national_id", request.NationalId).Msg("Successfully retrieved passenger by ID")
	return ctx.Status(http.StatusOK).JSON(passenger)
}

func (c *PassengerController) GetPassengerByPNR(ctx *fiber.Ctx) error {
	var request entities.GetPassengerByPnrRequest
	if err := ctx.QueryParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	passenger, err := c.service.GetPassengerByPNR(request)
	if err != nil {
		log.Error().Err(err).Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Error getting passenger by PNR")
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Passenger not found"})
	}
	log.Info().Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Successfully retrieved passenger by PNR")
	return ctx.Status(http.StatusOK).JSON(passenger)
}

func (c *PassengerController) OnlineCheckInPassenger(ctx *fiber.Ctx) error {
	var request entities.OnlineCheckInRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := c.service.OnlineCheckInPassenger(request)
	if err != nil {
		log.Error().Err(err).Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Error checking in passenger")
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Passenger not found or check-in failed"})
	}
	log.Info().Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Successfully checked in passenger")
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Check-in successful"})
}

func (c *PassengerController) GetPassengersBySpecificFlight(ctx *fiber.Ctx) error {
	var request entities.GetPassengersBySpecificFlightRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query parameters")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
	}

	passengers, err := c.service.GetPassengersBySpecificFlight(request)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error getting passengers by specific flight")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error retrieving passengers"})
	}

	log.Info().Str("flight_number", request.FlightNumber).Msg("Successfully retrieved passengers by specific flight")
	return ctx.Status(http.StatusOK).JSON(passengers)
}

func (c *PassengerController) CreatePassenger(ctx *fiber.Ctx) error {
	var request entities.CreatePassengerRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := c.service.CreatePassenger(request)
	if err != nil {
		log.Error().Err(err).Str("national_id", request.Passenger.NationalId).Msg("Error creating passenger")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating passenger"})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Passenger created successfully"})
}

func (c *PassengerController) GetAllPassengers(ctx *fiber.Ctx) error {
	passengers, err := c.service.GetAllPassengers()
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving all passengers")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error retrieving all passengers"})
	}
	return ctx.Status(http.StatusOK).JSON(passengers)
}

func (c *PassengerController) EmployeeCheckInPassenger(ctx *fiber.Ctx) error {
	var request entities.EmployeeCheckInRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	passenger, err := c.service.EmployeeCheckInPassenger(request)
	if err != nil {
		log.Error().Err(err).
			Str("national_id", request.NationalId).
			Str("destination_airport", request.DestinationAirport).
			Msg("Error checking in passenger")
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Passenger not found or check-in failed"})
	}

	log.Info().
		Str("national_id", request.NationalId).
		Str("destination_airport", request.DestinationAirport).
		Msg("Successfully checked in passenger")
	return ctx.Status(http.StatusOK).JSON(passenger)
}
