package app

import (
	"ams-service/config"
	"ams-service/infrastructure/api/controllers"
	"ams-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterEmployeeRoutes(router *gin.Engine, employeeController *controllers.EmployeeController) {
	employeeRoute := router.Group("/employee")
	{
		employeeRoute.GET("/:id", employeeController.GetEmployeeByID)
		employeeRoute.POST("/register", employeeController.RegisterEmployee)
		employeeRoute.POST("/login", employeeController.LoginEmployee)
		employeeRoute.Use(middlewares.AuthMiddleware(config.JWTSecretKey))
		//t
	}
}
