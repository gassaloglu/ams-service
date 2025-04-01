package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"
	"ams-service/internal/adapters/primary/rest/middlewares"
	"ams-service/internal/config"
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
