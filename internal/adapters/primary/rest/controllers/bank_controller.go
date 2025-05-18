package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type BankController struct {
	service primary.BankService
}

func NewBankController(service primary.BankService) *BankController {
	return &BankController{service: service}
}

func (c *BankController) CreateCreditCard(ctx *fiber.Ctx) error {
	if utils.IsBatchRequest(ctx) {
		var requests []entities.CreateCreditCardRequest
		if err := ctx.BodyParser(&requests); err != nil {
			log.Error().Err(err).Msg("Error binding JSON")
			return fiber.NewError(http.StatusBadRequest, "Invalid request")
		}

		err := c.service.CreateAllCreditCards(requests)
		if err != nil {
			log.Error().Err(err).Msg("Error adding credit cards")
			return fiber.NewError(http.StatusInternalServerError, "Server error")
		}

		return ctx.SendStatus(http.StatusOK)
	} else {
		var request entities.CreateCreditCardRequest
		if err := ctx.BodyParser(&request); err != nil {
			log.Error().Err(err).Msg("Error binding JSON")
			return fiber.NewError(http.StatusBadRequest, "Invalid request")
		}

		card, err := c.service.CreateCreditCard(&request)
		if err != nil {
			log.Error().Err(err).Msg("Error adding credit card")
			return fiber.NewError(http.StatusInternalServerError, "Server error")
		}

		log.Info().Uint("id", card.ID).Msg("Credit card added successfully")

		return ctx.SendStatus(http.StatusOK)
	}
}
