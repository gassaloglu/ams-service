package app

import (
	"ams-service/infrastructure/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	userRoute := router.Group("/user")
	{
		userRoute.POST("/register", userController.RegisterUser)
		userRoute.POST("/login", userController.LoginUser)
	}
}
