package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

/* ADAPTER - HANDLER */

type PlaneController struct {
	service primary.PlaneService
}

func NewPlaneController(service primary.PlaneService) *PlaneController {
	return &PlaneController{service: service}
}

func (c *PlaneController) GetAllPlanes(ctx *fiber.Ctx) error {
	planes, err := c.service.GetAllPlanes()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all planes")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting all planes"})
	}
	return ctx.Status(http.StatusOK).JSON(planes)
}

func (c *PlaneController) AddPlane(ctx *fiber.Ctx) error {
	var request entities.AddPlaneRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Error binding JSON"})
	}

	err := c.service.AddPlane(request)
	if err != nil {
		log.Error().Err(err).Msg("Error adding plane")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error adding plane"})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Plane added successfully"})
}

func (c *PlaneController) SetPlaneStatus(ctx *fiber.Ctx) error {
	var request entities.SetPlaneStatusRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Error binding JSON"})
	}

	err := c.service.SetPlaneStatus(request)
	if err != nil {
		log.Error().Err(err).Msg("Error setting plane status")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error setting plane status"})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Plane status updated successfully"})
}

func (c *PlaneController) GetPlaneByRegistration(ctx *fiber.Ctx) error {
	var request entities.GetPlaneByRegistrationRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Error binding query"})
	}

	plane, err := c.service.GetPlaneByRegistration(request)
	if err != nil {
		log.Error().Err(err).Str("registration", request.PlaneRegistration).Msg("Error getting plane by registration")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting plane by registration"})
	}
	return ctx.Status(http.StatusOK).JSON(plane)
}

func (c *PlaneController) GetPlaneByFlightNumber(ctx *fiber.Ctx) error {
	var request entities.GetPlaneByFlightNumberRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "TODO: Error binding query"})
	}

	plane, err := c.service.GetPlaneByFlightNumber(request)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error getting plane by flight number")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "TODO: Error getting plane by flight number"})
	}
	return ctx.Status(http.StatusOK).JSON(plane)
}

func (c *PlaneController) GetPlaneByLocation(ctx *fiber.Ctx) error {
	var request entities.GetPlaneByLocationRequest
	if err := ctx.QueryParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "TODO: Error binding query"})
	}

	planes, err := c.service.GetPlaneByLocation(request)
	if err != nil {
		log.Error().Err(err).Str("location", request.Location).Msg("Error getting planes by location")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "TODO: Error getting planes by location"})
	}
	return ctx.Status(http.StatusOK).JSON(planes)
}
