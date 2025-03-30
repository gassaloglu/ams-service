package controllers

import (
	"ams-service/core/entities"
	"ams-service/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "TODO: Error binding JSON"})
		return
	}

	err := c.service.RegisterUser(user)
	if err != nil {
		log.Error().Err(err).Str("username", user.Username).Msg("Error registering user")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "TODO: Registration failed"})
		return
	}
	log.Info().Str("username", user.Username).Msg("Successfully registered user")
	ctx.JSON(http.StatusOK, gin.H{"message": "TODO: Registration successful"})
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	var loginRequest entities.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := c.service.LoginUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "user": user})
}
