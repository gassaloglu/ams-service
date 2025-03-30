package controllers

import (
	"ams-service/core/entities"
	"ams-service/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var PASSENGER_LOG_PREFIX string = "passenger_controller.go"

type PassengerController struct {
	service *services.PassengerService
}

func NewPassengerController(service *services.PassengerService) *PassengerController {
	return &PassengerController{service: service}
}

func (c *PassengerController) GetPassengerByID(ctx *gin.Context) {
	var request entities.GetPassengerByIdRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passenger, err := c.service.GetPassengerByID(request)
	if err != nil {
		log.Error().Err(err).Str("national_id", request.NationalId).Msg("Error getting passenger by ID")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "TODO: Passenger not found"})
		return
	}
	log.Info().Str("national_id", request.NationalId).Msg("Successfully retrieved passenger by ID")
	ctx.JSON(http.StatusOK, passenger)
}

func (c *PassengerController) GetPassengerByPNR(ctx *gin.Context) {
	var request entities.GetPassengerByPnrRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passenger, err := c.service.GetPassengerByPNR(request)
	if err != nil {
		log.Error().Err(err).Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Error getting passenger by PNR")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "TODO: Passenger not found"})
		return
	}
	log.Info().Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Successfully retrieved passenger by PNR")
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
		log.Error().Err(err).Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Error checking in passenger")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "TODO: Passenger not found or check-in failed"})
		return
	}
	log.Info().Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Successfully checked in passenger")
	ctx.JSON(http.StatusOK, gin.H{"message": "TODO: Check-in successful"})
}
