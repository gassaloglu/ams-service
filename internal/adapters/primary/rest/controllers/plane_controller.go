package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type PlaneController struct {
	service primary.PlaneService
}

func NewPlaneController(service primary.PlaneService) *PlaneController {
	return &PlaneController{service: service}
}

func (c *PlaneController) GetAllPlanes(ctx *fiber.Ctx) error {
	var query entities.GetAllPlanesRequest
	if err := ctx.QueryParser(&query); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		return fiber.NewError(http.StatusBadRequest, "Error binding query")
	}

	planes, err := c.service.GetAllPlanes(query)
	if err != nil {
		log.Error().Err(err).Msg("Error getting all planes")
		return fiber.NewError(http.StatusInternalServerError, "Error getting all planes")
	}

	log.Info().Msg("Successfully retrieved all planes")

	return ctx.Status(http.StatusOK).JSON(planes)
}

func (c *PlaneController) AddPlane(ctx *fiber.Ctx) error {
	var request entities.AddPlaneRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return fiber.NewError(http.StatusBadRequest, "Error binding query")
	}

	err := c.service.AddPlane(request)
	if err != nil {
		log.Error().Err(err).Msg("Error adding plane")
		return fiber.NewError(http.StatusInternalServerError, "Error adding plane")
	}

	log.Info().Msgf("Successfully added plane(s)")

	return ctx.SendStatus(http.StatusCreated)
}

func (c *PlaneController) SetPlaneStatus(ctx *fiber.Ctx) error {
	var request entities.SetPlaneStatusRequest

	if err := ctx.ParamsParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return fiber.NewError(http.StatusBadRequest, "Error binding query")
	}

	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return fiber.NewError(http.StatusBadRequest, "Error binding query")
	}

	err := c.service.SetPlaneStatus(request)
	if err != nil {
		log.Error().Err(err).Msg("Error setting plane status")
		return fiber.NewError(http.StatusInternalServerError, "Error setting plane status")
	}

	log.Info().Msgf("Successfully set plane status")

	return ctx.SendStatus(http.StatusOK)
}

func (c *PlaneController) GetPlaneByRegistration(ctx *fiber.Ctx) error {
	var request entities.GetPlaneByRegistrationRequest
	if err := ctx.ParamsParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		return fiber.NewError(http.StatusBadRequest, "Error binding query")
	}

	plane, err := c.service.GetPlaneByRegistration(request)
	if err != nil {
		log.Error().Err(err).Str("registration", request.Registration).Msg("Error getting plane by registration")
		return fiber.NewError(http.StatusInternalServerError, "Error getting plane by registration")
	}

	log.Info().Str("registration", request.Registration).Msg("Successfully retrieved plane by registration")

	return ctx.Status(http.StatusOK).JSON(plane)
}
