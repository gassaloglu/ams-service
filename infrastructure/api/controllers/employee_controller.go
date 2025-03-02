package controllers

import (
	"ams-service/core/entities"
	"ams-service/core/services"
	"ams-service/middlewares"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var EMPLOYEE_LOG_PREFIX string = "employee_controller.go"

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
		middlewares.LogError(EMPLOYEE_LOG_PREFIX + " - Error converting ID: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	request := entities.GetEmployeeByIdRequest{ID: uint(id)}
	employee, err := c.service.GetEmployeeByID(request)
	if err != nil {
		middlewares.LogError(EMPLOYEE_LOG_PREFIX + " - Error getting employee by ID: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	ctx.JSON(http.StatusOK, employee)
}
