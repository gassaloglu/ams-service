package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var PLANE_LOG_PREFIX string = "plane_controller.go"

/* ADAPTER - HANDLER */

type PlaneController struct {
	service *services.PlaneService
}

func NewPlaneController(service *services.PlaneService) *PlaneController {
	return &PlaneController{service: service}
}

func (c *PlaneController) GetAllPlanes(ctx *gin.Context) {
	planes, err := c.service.GetAllPlanes()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all planes")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "TODO: Error getting all planes"})
		return
	}
	ctx.JSON(http.StatusOK, planes)
}

func (c *PlaneController) AddPlane(ctx *gin.Context) {
	var request entities.AddPlaneRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "TODO: Error binding JSON"})
		return
	}

	err := c.service.AddPlane(request)
	if err != nil {
		log.Error().Err(err).Msg("Error adding plane")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "TODO: Error adding plane"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Plane added successfully"})
}

func (c *PlaneController) SetPlaneStatus(ctx *gin.Context) {
	var request entities.SetPlaneStatusRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "TODO: Error binding JSON"})
		return
	}

	err := c.service.SetPlaneStatus(request)
	if err != nil {
		log.Error().Err(err).Msg("Error setting plane status")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "TODO: Error setting plane status"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Plane status updated successfully"})
}

func (c *PlaneController) GetPlaneByRegistration(ctx *gin.Context) {
	var request entities.GetPlaneByRegistrationRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error binding query"})
		return
	}

	plane, err := c.service.GetPlaneByRegistration(request)
	if err != nil {
		log.Error().Err(err).Str("registration", request.PlaneRegistration).Msg("Error getting plane by registration")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting plane by registration"})
		return
	}
	ctx.JSON(http.StatusOK, plane)
}

func (c *PlaneController) GetPlaneByFlightNumber(ctx *gin.Context) {
	var request entities.GetPlaneByFlightNumberRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "TODO: Error binding query"})
		return
	}

	plane, err := c.service.GetPlaneByFlightNumber(request)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error getting plane by flight number")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "TODO: Error getting plane by flight number"})
		return
	}
	ctx.JSON(http.StatusOK, plane)
}

func (c *PlaneController) GetPlaneByLocation(ctx *gin.Context) {
	var request entities.GetPlaneByLocationRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Error().Err(err).Msg("Error binding query")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "TODO: Error binding query"})
		return
	}

	planes, err := c.service.GetPlaneByLocation(request)
	if err != nil {
		log.Error().Err(err).Str("location", request.Location).Msg("Error getting planes by location")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "TODO: Error getting planes by location"})
		return
	}
	ctx.JSON(http.StatusOK, planes)
}
