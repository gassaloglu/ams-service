package controllers

import (
	"ams-service/core/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* ADAPTER - HANDLER */

type PassengerController struct {
	service *services.PassengerService
}

func NewPassengerController(service *services.PassengerService) *PassengerController {
	return &PassengerController{service: service}
}

func (c *PassengerController) GetPassengerByID(ctx *gin.Context) {
	passengerID := ctx.Param("id")
	passenger, err := c.service.GetPassengerByID(passengerID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Passenger not found"})
		return
	}
	ctx.JSON(http.StatusOK, passenger)
}

func (c *PassengerController) OnlineCheckInPassenger(ctx *gin.Context) {
	pnr := ctx.Param("pnr")
	surname := ctx.Param("surname")

	err := c.service.OnlineCheckInPassenger(pnr, surname)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Passenger not found or check-in failed"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Check-in successful"})
}
