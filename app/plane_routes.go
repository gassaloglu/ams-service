package app

import (
	"ams-service/infrastructure/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPlaneRoutes(router *gin.Engine, planeController *controllers.PlaneController) {
	planeRoute := router.Group("/plane")
	{
		planeRoute.GET("/all", planeController.GetAllPlanes)
		planeRoute.POST("/add", planeController.AddPlane)
		planeRoute.PUT("/status", planeController.SetPlaneStatus)
		planeRoute.GET("/registration", planeController.GetPlaneByRegistration)
		planeRoute.GET("/flightnumber", planeController.GetPlaneByFlightNumber)
		planeRoute.GET("/location", planeController.GetPlaneByLocation)
	}
}
