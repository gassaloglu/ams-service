package controllers

import (
	"ams-service/core/entities"
	"ams-service/core/services"
	"ams-service/middlewares"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const USER_LOG_PREFIX string = "user_controller.go"

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.RegisterUser(user)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering user: %v", USER_LOG_PREFIX, err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully registered user: %v", USER_LOG_PREFIX, user))
	ctx.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
