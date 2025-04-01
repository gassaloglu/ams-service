package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/core/services"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type EmployeeController struct {
	service *services.EmployeeService
}

func NewEmployeeController(service *services.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (c *EmployeeController) GetEmployeeByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Error().Err(err).Str("id", idParam).Msg("Error converting ID")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	request := entities.GetEmployeeByIdRequest{ID: uint(id)}
	employee, err := c.service.GetEmployeeByID(request)
	if err != nil {
		log.Error().Err(err).Uint("id", uint(id)).Msg("Error getting employee by ID")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	ctx.JSON(http.StatusOK, employee)
}

func (c *EmployeeController) RegisterEmployee(ctx *gin.Context) {
	var request entities.RegisterEmployeeRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.service.RegisterEmployee(request)
	if err != nil {
		log.Error().Err(err).Msg("Error registering employee")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Employee registered successfully"})
}

func (c *EmployeeController) LoginEmployee(ctx *gin.Context) {
	var loginRequest entities.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	employee, token, err := c.service.LoginEmployee(context.Background(), loginRequest.Username, loginRequest.Password)
	if err != nil {
		log.Error().Err(err).Msg("Error logging in employee")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid employee ID or password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "employee": employee})
}
