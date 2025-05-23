package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/utils"

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
	if utils.IsBatchRequest(ctx) {
		var requests []entities.RegisterUserRequest
		if err := ctx.BodyParser(&requests); err != nil {
			log.Error().Err(err).Msg("Error binding register request")
			return fiber.NewError(fiber.StatusBadRequest, "Malformed register request body")
		}

		err := c.service.RegisterAll(requests)
		if err != nil {
			log.Error().Err(err).Msg("Error registering user")
			return fiber.NewError(fiber.StatusInternalServerError, "Error registering user")
		}

		return ctx.SendStatus(fiber.StatusCreated)
	} else {
		var request entities.RegisterUserRequest
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

		response := &entities.RegisterUserResponse{Token: token}
		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}

func (c *UserController) LoginUser(ctx *fiber.Ctx) error {
	var request entities.LoginUserRequest
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

	response := &entities.LoginUserResponse{Token: token}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := c.service.GetAllUsers()
	if err != nil {
		log.Error().Err(err).Msg("Error getting users")
		return fiber.NewError(fiber.StatusInternalServerError, "Error getting users")
	}

	return ctx.Status(fiber.StatusCreated).JSON(users)
}
