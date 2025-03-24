package app

import (
	"ams-service/config"
	"ams-service/infrastructure/api/controllers"
	"ams-service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	userRoute := router.Group("/user")
	{
		userRoute.POST("/register", userController.RegisterUser)
		userRoute.POST("/login", userController.LoginUser)
		userRoute.Use(middlewares.AuthMiddleware(config.JWTSecretKey))
	}
}
