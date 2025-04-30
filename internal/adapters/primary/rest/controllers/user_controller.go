package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
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
	var request entities.UserRegisterRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding register request")
		return fiber.NewError(fiber.StatusBadRequest, "Malformed register request body")
	}

	token, err := c.service.Register(&request)
	if err != nil {
		log.Error().Err(err).Msg("Error registering user")
		return fiber.NewError(fiber.StatusInternalServerError, "Error registering user")
	}

	log.Info().Str("email", request.Email).Msg("Successfully registered user")

	response := &entities.UserRegisterResponse{Token: token}
	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *UserController) LoginUser(ctx *fiber.Ctx) error {
	var request entities.UserLoginRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding register request")
		return fiber.NewError(fiber.StatusBadRequest, "Malformed register request body")
	}

	token, err := c.service.Login(request.Email, request.Password)
	if err != nil {
		log.Error().Err(err).Msg("Error logging user in")
		return fiber.NewError(fiber.StatusInternalServerError, "Error logging user in")
	}

	log.Info().Str("email", request.Email).Msg("Login successful")
	
	response := &entities.UserLoginResponse{Token: token}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
