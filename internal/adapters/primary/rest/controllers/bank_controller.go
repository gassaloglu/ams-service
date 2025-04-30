package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
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

func (c *BankController) AddCreditCard(ctx *fiber.Ctx) error {
	var card entities.CreditCard
	if err := ctx.BodyParser(&card); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := c.service.AddCreditCard(card)
	if err != nil {
		log.Error().Err(err).Msg("Error adding credit card")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Credit card added successfully"})
}

func (c *BankController) GetAllCreditCards(ctx *fiber.Ctx) error {
	cards, err := c.service.GetAllCreditCards()
	if err != nil {
		log.Error().Err(err).Msg("Error getting credit cards")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}
	return ctx.Status(http.StatusOK).JSON(cards)
}

func (c *BankController) Pay(ctx *fiber.Ctx) error {
	var request entities.PaymentRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := c.service.Pay(request)
	if err != nil {
		log.Error().Err(err).Msg("Error processing payment")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Payment processed successfully"})
}

func (c *BankController) Refund(ctx *fiber.Ctx) error {
	var request entities.Refund
	if err := ctx.BodyParser(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := c.service.Refund(request)
	if err != nil {
		log.Error().Err(err).Msg("Error processing refund")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Refund processed successfully"})
}
