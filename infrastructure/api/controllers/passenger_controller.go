package controllers

import (
	"ams-service/core/entities"
	"ams-service/core/services"
	"ams-service/middlewares"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var LOG_PREFIX string = "passenger_controller.go"

/* ADAPTER - HANDLER */

type PassengerController struct {
	service *services.PassengerService
}

func NewPassengerController(service *services.PassengerService) *PassengerController {
	return &PassengerController{service: service}
}

func (c *PassengerController) GetPassengerByID(ctx *gin.Context) {
	passengerID := ctx.Param("id")
	if passengerID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Passenger ID is required"})
		return
	}

	passenger, err := c.service.GetPassengerByID(passengerID)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting passenger by ID %s: %v", LOG_PREFIX, passengerID, err))
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Passenger not found"})
		return
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved passenger by ID %s", LOG_PREFIX, passengerID))
	ctx.JSON(http.StatusOK, passenger)
}

func (c *PassengerController) OnlineCheckInPassenger(ctx *gin.Context) {
	var request entities.OnlineCheckInRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.OnlineCheckInPassenger(request.PNR, request.Surname)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error checking in passenger with PNR %s and surname %s: %v", LOG_PREFIX, request.PNR, request.Surname, err))
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Passenger not found or check-in failed"})
		return
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully checked in passenger with PNR %s and surname %s", LOG_PREFIX, request.PNR, request.Surname))
	ctx.JSON(http.StatusOK, gin.H{"message": "Check-in successful"})
}
