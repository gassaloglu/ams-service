package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/utils"
	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type FlightController struct {
	service primary.FlightService
}

func NewFlightController(service primary.FlightService) *FlightController {
	return &FlightController{service: service}
}

func (c *FlightController) GetFlightById(ctx *fiber.Ctx) error {
	var request entities.GetFlightByIdRequest
	if err := ctx.ParamsParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid URI params")
	}

	flight, err := c.service.FindById(&request)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).Msg("Flight not found")
		return fiber.NewError(fiber.StatusNotFound, "Flight not found")
	} else if err != nil {
		log.Error().Err(err).Msg("Error finding flight by id")
		return fiber.NewError(fiber.StatusInternalServerError, "Error finding flight by id")
	}

	return ctx.Status(http.StatusOK).JSON(flight)
}

func (c *FlightController) GetAllFlights(ctx *fiber.Ctx) error {
	var request entities.GetAllFlightsRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query params")
	}

	flights, err := c.service.FindAll(&request)
	if err != nil {
		log.Error().Err(err).Msg("Error getting all flights")
		return fiber.NewError(http.StatusInternalServerError, "Error getting all flights")
	}

	return ctx.Status(http.StatusOK).JSON(flights)
}

func (c *FlightController) GetAllActiveFlights(ctx *fiber.Ctx) error {
	var request entities.GetAllFlightsRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query params")
	}

	flights, err := c.service.FindAllActive(&request)
	if err != nil {
		log.Error().Err(err).Msg("Error getting all flights")
		return fiber.NewError(http.StatusInternalServerError, "Error getting all flights")
	}

	return ctx.Status(http.StatusOK).JSON(flights)
}

func (c *FlightController) CreateFlight(ctx *fiber.Ctx) error {
	if utils.IsBatchRequest(ctx) {
		var requests []entities.CreateFlightRequest
		if err := ctx.BodyParser(&requests); err != nil {
			log.Error().Err(err).Msg("Error binding JSON")
			return fiber.NewError(http.StatusBadRequest, "Invalid request payload")
		}

		// Call the service to add the flight
		err := c.service.CreateAll(requests)
		if err != nil {
			log.Error().Err(err).Msg("Error adding flights")
			return fiber.NewError(fiber.StatusInternalServerError, "Error adding flights")
		}

		log.Info().Msg("Successfully added flights")
		return ctx.SendStatus(http.StatusCreated)
	} else {
		// Parse the request body
		var request entities.CreateFlightRequest
		if err := ctx.BodyParser(&request); err != nil {
			log.Error().Err(err).Msg("Error binding JSON")
			return fiber.NewError(http.StatusBadRequest, "Invalid request payload")
		}

		// Call the service to add the flight
		err := c.service.Create(&request)
		if err != nil {
			log.Error().Err(err).Msg("Error adding flight")
			return fiber.NewError(fiber.StatusInternalServerError, "Error adding flight")
		}

		log.Info().Msg("Successfully added flight")
		return ctx.SendStatus(http.StatusCreated)
	}
}

func (c *FlightController) GetSeatsByFlightId(ctx *fiber.Ctx) error {
	var request entities.GetSeatsByFlightIdRequest
	if err := ctx.ParamsParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid URI params")
	}

	seats, err := c.service.FindSeatsByFlightId(&request)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).Msg("Flight not found")
		return fiber.NewError(fiber.StatusNotFound, "Flight not found")
	} else if err != nil {
		log.Error().Err(err).Msg("Error finding flight by id")
		return fiber.NewError(fiber.StatusInternalServerError, "Error finding flight by id")
	}

	return ctx.Status(http.StatusOK).JSON(seats)
}
