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

	planes, err := c.service.FindAll(&query)
	if err != nil {
		log.Error().Err(err).Msg("Error getting all planes")
		return fiber.NewError(http.StatusInternalServerError, "Error getting all planes")
	}

	log.Info().Msg("Successfully retrieved all planes")

	return ctx.Status(http.StatusOK).JSON(planes)
}

func (c *PlaneController) CreatePlane(ctx *fiber.Ctx) error {
	var request entities.CreatePlaneRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return fiber.NewError(http.StatusBadRequest, "Error binding query")
	}

	plane, err := c.service.Create(&request)
	if err != nil {
		log.Error().Err(err).Msg("Error adding plane")
		return fiber.NewError(http.StatusInternalServerError, "Error adding plane")
	}

	log.Info().Msgf("Successfully added plane(s)")

	return ctx.Status(http.StatusCreated).JSON(plane)
}
