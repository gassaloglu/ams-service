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

func (c *UserController) RegisterUser(ctx *fiber.Ctx) error {
	var user entities.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Error binding JSON",
		})
	}

	err := c.service.RegisterUser(user)
	if err != nil {
		log.Error().Err(err).Str("username", user.Username).Msg("Error registering user")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Registration failed",
		})
	}

	log.Info().Str("username", user.Username).Msg("Successfully registered user")
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Registration successful",
	})
}

func (c *UserController) LoginUser(ctx *fiber.Ctx) error {
	var loginRequest entities.LoginRequest
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, token, err := c.service.LoginUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
		"user":    user,
	})
}
