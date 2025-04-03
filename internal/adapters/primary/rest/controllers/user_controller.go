package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type UserController struct {
	service primary.UserService
}

func NewUserController(service primary.UserService) *UserController {
	return &UserController{service: service}
}

// Helper function to handle errors and send JSON responses
func respondWithError(ctx *fiber.Ctx, status int, message string, err error) error {
	if err != nil {
		log.Error().Err(err).Msg(message)
	}
	return ctx.Status(status).JSON(fiber.Map{"error": message})
}

func (c *UserController) RegisterUser(ctx *fiber.Ctx) error {
	var user entities.User
	if err := ctx.BodyParser(&user); err != nil {
		return respondWithError(ctx, http.StatusBadRequest, "Error binding JSON", err)
	}

	if err := c.service.RegisterUser(user); err != nil {
		return respondWithError(ctx, http.StatusInternalServerError, "Registration failed", err)
	}

	log.Info().Str("username", user.Username).Msg("Successfully registered user")
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Registration successful"})
}

func (c *UserController) LoginUser(ctx *fiber.Ctx) error {
	var loginRequest entities.LoginRequest
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return respondWithError(ctx, http.StatusBadRequest, "Error binding JSON", err)
	}

	user, token, err := c.service.LoginUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return respondWithError(ctx, http.StatusUnauthorized, "Invalid username or password", err)
	}

	log.Info().Str("username", loginRequest.Username).Msg("Login successful")
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
		"user":    user,
	})
}
