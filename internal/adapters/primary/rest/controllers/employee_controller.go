package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type EmployeeController struct {
	service primary.EmployeeService
}

func NewEmployeeController(service primary.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (c *EmployeeController) GetEmployeeByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Error().Err(err).Str("id", idParam).Msg("Error converting ID")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	request := entities.GetEmployeeByIdRequest{ID: uint(id)}
	employee, err := c.service.GetEmployeeByID(request)
	if err != nil {
		log.Error().Err(err).Uint("id", uint(id)).Msg("Error getting employee by ID")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}
	return ctx.Status(http.StatusOK).JSON(employee)
}

func (c *EmployeeController) RegisterEmployee(ctx *fiber.Ctx) error {
	var request entities.RegisterEmployeeRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := c.service.RegisterEmployee(request)
	if err != nil {
		log.Error().Err(err).Msg("Error registering employee")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Employee registered successfully"})
}

func (c *EmployeeController) LoginEmployee(ctx *fiber.Ctx) error {
	var loginRequest entities.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	employee, token, err := c.service.LoginEmployee(loginRequest.Username, loginRequest.Password)
	if err != nil {
		log.Error().Err(err).Msg("Error logging in employee")
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid employee ID or password"})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Login successful", "token": token, "employee": employee})
}
