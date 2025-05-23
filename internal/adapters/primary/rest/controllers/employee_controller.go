package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type EmployeeController struct {
	service primary.EmployeeService
}

func NewEmployeeController(service primary.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (c *EmployeeController) GetEmployees(ctx *fiber.Ctx) error {
	employees, err := c.service.FindAll()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all employees")
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(employees)
}

func (c *EmployeeController) RegisterEmployee(ctx *fiber.Ctx) error {
	if utils.IsBatchRequest(ctx) {
		var requests []entities.RegisterEmployeeRequest
		if err := ctx.BodyParser(&requests); err != nil {
			log.Error().Err(err).Msg("Error binding JSON")
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err := c.service.RegisterAll(requests)
		if err != nil {
			log.Error().Err(err).Msg("Error registering employees")
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return ctx.SendStatus(http.StatusCreated)
	} else {
		var request entities.RegisterEmployeeRequest
		if err := ctx.BodyParser(&request); err != nil {
			log.Error().Err(err).Msg("Error binding JSON")
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		token, err := c.service.Register(&request)
		if err != nil {
			log.Error().Err(err).Msg("Error registering employee")
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		response := entities.RegisterEmployeeResponse{Token: token}
		return ctx.Status(http.StatusCreated).JSON(response)
	}
}

func (c *EmployeeController) LoginEmployee(ctx *fiber.Ctx) error {
	var request entities.LoginEmployeeRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := c.service.Login(&request)
	if err != nil {
		log.Error().Err(err).Msg("Error logging employee in")
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON(response)
}
