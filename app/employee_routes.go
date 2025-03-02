package app

import (
	"ams-service/infrastructure/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterEmployeeRoutes(router *gin.Engine, employeeController *controllers.EmployeeController) {
	employeeRoute := router.Group("/employee")
	{
		// TODO: url binding will be changed
		// Parameter structure will be used instead of path
		employeeRoute.GET("/:id", employeeController.GetEmployeeByID)
		employeeRoute.POST("/register", employeeController.RegisterEmployee)
	}
}
