package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"
	"ams-service/internal/adapters/primary/rest/middlewares"
	"ams-service/internal/config"
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
