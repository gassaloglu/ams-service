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

func (c *PassengerController) GetPassengersBySpecificFlight(ctx *gin.Context) {
	var request entities.GetPassengersBySpecificFlightRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query parameters")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	passengers, err := c.service.GetPassengersBySpecificFlight(request)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error getting passengers by specific flight")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving passengers"})
		return
	}

	log.Info().Str("flight_number", request.FlightNumber).Msg("Successfully retrieved passengers by specific flight")
	ctx.JSON(http.StatusOK, passengers)
}

func (c *PassengerController) CreatePassenger(ctx *gin.Context) {
	var request entities.CreatePassengerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.service.CreatePassenger(request)
	if err != nil {
		log.Error().Err(err).Str("national_id", request.NationalId).Msg("Error creating passenger")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating passenger"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Passenger created successfully"})
}

func (c *PassengerController) GetAllPassengers(ctx *gin.Context) {
	passengers, err := c.service.GetAllPassengers()
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving all passengers")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving all passengers"})
		return
	}
	ctx.JSON(http.StatusOK, passengers)
}
