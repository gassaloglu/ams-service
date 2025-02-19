package controllers

import (
	"ams-service/core/services"

	"github.com/gin-gonic/gin"
)

var PLANE_LOG_PREFIX string = "plane_controller.go"

/* ADAPTER - HANDLER */

type PlaneController struct {
	service *services.PlaneService
}

func NewPlaneController(service *services.PlaneService) *PlaneController {
	return &PlaneController{service: service}
}

func (c *PlaneController) GetAllplanes(ctx *gin.Context) {
	// Will be added
}

func (c *PlaneController) AddPlane(ctx *gin.Context) {
	// Will be added
}

func (c *PlaneController) SetPlaneStatus(ctx *gin.Context) {
	// Will be added
}

func (c *PlaneController) GetPlaneByRegistration(ctx *gin.Context) {
	// Will be added
}

func (c *PlaneController) GetPlaneByFlightNumber(ctx *gin.Context) {
	// Will be added
}

func (c *PlaneController) GetPlaneByLocation(ctx *gin.Context) {
	// Will be added
}
