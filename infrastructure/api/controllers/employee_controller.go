package controllers

import (
	"ams-service/core/entities"
	"ams-service/core/services"
	"ams-service/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

var EMPLOYEE_LOG_PREFIX string = "employee_controller.go"

/* ADAPTER - HANDLER */

type EmployeeController struct {
	service *services.EmployeeService
}

func NewEmployeeController(service *services.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (c *EmployeeController) RegisterEmployee(ctx *gin.Context) {
	var request entities.RegisterEmployeeRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		middlewares.LogError(EMPLOYEE_LOG_PREFIX + " - Error binding JSON: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.RegisterEmployee(request)
	if err != nil {
		middlewares.LogError(EMPLOYEE_LOG_PREFIX + " - Error registering employee: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Employee registered successfully"})
}

func (c *EmployeeController) GetEmployeeByNationalID(ctx *gin.Context) {
	var request entities.GetEmployeeByNationalIDRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		middlewares.LogError(EMPLOYEE_LOG_PREFIX + " - Error binding query: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := c.service.GetEmployeeByNationalID(request)
	if err != nil {
		middlewares.LogError(EMPLOYEE_LOG_PREFIX + " - Error getting employee by national ID: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, employee)
}

func (c *EmployeeController) GetEmployeeByID(ctx *gin.Context) {
	var request entities.GetEmployeeByIDRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		middlewares.LogError(EMPLOYEE_LOG_PREFIX + " - Error binding query: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := c.service.GetEmployeeByID(request)
	if err != nil {
		middlewares.LogError(EMPLOYEE_LOG_PREFIX + " - Error getting employee by ID: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, employee)
}
