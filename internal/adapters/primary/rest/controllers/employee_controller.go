package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
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

func (c *EmployeeController) GetEmployeeByID(ctx *fiber.Ctx) error {
	employeeID := ctx.Query("employee_id")
	if employeeID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Employee ID is required",
		})
	}

	request := entities.GetEmployeeByIdRequest{EmployeeID: employeeID}
	employee, err := c.service.GetEmployeeByID(request)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Employee not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(employee)
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
	var loginRequest entities.LoginEmployeeRequest
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	employee, token, err := c.service.LoginEmployee(loginRequest.EmployeeID, loginRequest.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid employee ID or password",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message":  "Login successful",
		"token":    token,
		"employee": employee,
	})
}
