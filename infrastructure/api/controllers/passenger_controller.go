package controllers

import (
	"ams-service/core/entities"
	"ams-service/core/services"
	"ams-service/middlewares"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var PASSENGER_LOG_PREFIX string = "passenger_controller.go"

/* ADAPTER - HANDLER */

type PassengerController struct {
	service *services.PassengerService
}

func NewPassengerController(service *services.PassengerService) *PassengerController {
	return &PassengerController{service: service}
}

func (c *PassengerController) GetPassengerByID(ctx *gin.Context) {
	var request entities.GetPassengerByIdRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passenger, err := c.service.GetPassengerByID(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting passenger by ID %s: %v", PASSENGER_LOG_PREFIX, request.NationalId, err))
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Passenger not found"})
		return
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved passenger by ID %s", PASSENGER_LOG_PREFIX, request.NationalId))
	ctx.JSON(http.StatusOK, passenger)
}

func (c *PassengerController) GetPassengerByPNR(ctx *gin.Context) {
	var request entities.GetPassengerByPnrRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passenger, err := c.service.GetPassengerByPNR(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting passenger by PNR %s and surname %s: %v", PASSENGER_LOG_PREFIX, request.PNR, request.Surname, err))
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Passenger not found"})
		return
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved passenger by PNR %s and surname %s", PASSENGER_LOG_PREFIX, request.PNR, request.Surname))
	ctx.JSON(http.StatusOK, passenger)
}

func (c *PassengerController) OnlineCheckInPassenger(ctx *gin.Context) {
	var request entities.OnlineCheckInRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.OnlineCheckInPassenger(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error checking in passenger with PNR %s and surname %s: %v", PASSENGER_LOG_PREFIX, request.PNR, request.Surname, err))
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Passenger not found or check-in failed"})
		return
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully checked in passenger with PNR %s and surname %s", PASSENGER_LOG_PREFIX, request.PNR, request.Surname))
	ctx.JSON(http.StatusOK, gin.H{"message": "Check-in successful"})
}
