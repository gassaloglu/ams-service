package controllers

import "ams-service/core/services"

var PLANE_LOG_PREFIX string = "plane_controller.go"

/* ADAPTER - HANDLER */

type PlaneController struct {
	service *services.PlaneService
}

func NewPlaneController(service *services.PlaneService) *PlaneController {
	return &PlaneController{service: service}
}
