package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/utils"
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
		log.Error().Err(err).Msg("Error binding query")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query params")
	}

	passenger, err := c.service.GetPassengerByID(request)
	if err != nil {
		log.Error().Err(err).Str("national_id", request.NationalId).Msg("Error getting passenger by ID")
		return fiber.NewError(fiber.StatusNotFound, "Passenger not found")
	}
	log.Info().Str("national_id", request.NationalId).Msg("Successfully retrieved passenger by ID")
	return ctx.Status(http.StatusOK).JSON(passenger)
}

func (c *PassengerController) GetPassengerByPNR(ctx *fiber.Ctx) error {
	var request entities.GetPassengerByPnrRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query params")
	}

	passenger, err := c.service.GetPassengerByPNR(request)
	if err != nil {
		log.Error().Err(err).Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Error getting passenger by PNR")
		return fiber.NewError(fiber.StatusNotFound, "Passenger not found")
	}
	log.Info().Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Successfully retrieved passenger by PNR")
	return ctx.Status(http.StatusOK).JSON(passenger)
}

func (c *PassengerController) OnlineCheckInPassenger(ctx *fiber.Ctx) error {
	var request entities.OnlineCheckInRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	err := c.service.OnlineCheckInPassenger(request)
	if err != nil {
		log.Error().Err(err).Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Error checking in passenger")
		return fiber.NewError(fiber.StatusNotFound, "Passenger not found or check-in failed")
	}
	log.Info().Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Successfully checked in passenger")
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Check-in successful"})
}

func (c *PassengerController) GetPassengersBySpecificFlight(ctx *fiber.Ctx) error {
	var request entities.GetPassengersBySpecificFlightRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query parameters")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	passengers, err := c.service.GetPassengersBySpecificFlight(request)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error getting passengers by specific flight")
		return fiber.NewError(fiber.StatusInternalServerError, "Error retrieving passengers")
	}

	log.Info().Str("flight_number", request.FlightNumber).Msg("Successfully retrieved passengers by specific flight")
	return ctx.Status(http.StatusOK).JSON(passengers)
}

func (c *PassengerController) CreatePassenger(ctx *fiber.Ctx) error {
	if utils.IsBatchRequest(ctx) {
		var requests []entities.CreatePassengerRequest
		if err := ctx.BodyParser(&requests); err != nil {
			log.Error().Err(err).Msg("Error binding JSON")
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
		}

		err := c.service.CreateAllPassengers(&requests)
		if err != nil {
			log.Error().Err(err).Msg("Error creating passengers")
			return fiber.NewError(fiber.StatusInternalServerError, "Error creating passenger")
		}

		return ctx.SendStatus(http.StatusCreated)
	} else {
		var request entities.CreatePassengerRequest
		if err := ctx.BodyParser(&request); err != nil {
			log.Error().Err(err).Msg("Error binding JSON")
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
		}

		_, err := c.service.CreatePassenger(&request)
		if err != nil {
			log.Error().Err(err).Msg("Error creating passenger")
			return fiber.NewError(fiber.StatusInternalServerError, "Error creating passenger")
		}

		return ctx.SendStatus(http.StatusCreated)
	}
}

func (c *PassengerController) GetAllPassengers(ctx *fiber.Ctx) error {
	passengers, err := c.service.GetAllPassengers()
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving all passengers")
		return fiber.NewError(fiber.StatusInternalServerError, "Error retrieving all passengers")
	}
	return ctx.Status(http.StatusOK).JSON(passengers)
}

func (c *PassengerController) EmployeeCheckInPassenger(ctx *fiber.Ctx) error {
	var request entities.EmployeeCheckInRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	passenger, err := c.service.EmployeeCheckInPassenger(request)
	if err != nil {
		log.Error().Err(err).
			Str("national_id", request.NationalId).
			Str("destination_airport", request.DestinationAirport).
			Msg("Error checking in passenger")
		return fiber.NewError(fiber.StatusNotFound, "Passenger not found or check-in failed")
	}

	log.Info().
		Str("national_id", request.NationalId).
		Str("destination_airport", request.DestinationAirport).
		Msg("Successfully checked in passenger")
	return ctx.Status(http.StatusOK).JSON(passenger)
}

func (c *PassengerController) CancelPassenger(ctx *fiber.Ctx) error {
	var request entities.CancelPassengerRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error parsing JSON body")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	err := c.service.CancelPassenger(request)
	if err != nil {
		log.Error().Err(err).Uint("passenger_id", request.PassengerID).Msg("Error canceling passenger")
		return fiber.NewError(fiber.StatusInternalServerError, "Error canceling passenger")
	}

	log.Info().Uint("passenger_id", request.PassengerID).Msg("Successfully canceled passenger")
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Passenger canceled successfully"})
}
