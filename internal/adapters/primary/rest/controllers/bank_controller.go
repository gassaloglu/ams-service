package controllers

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type BankController struct {
	service primary.BankService
}

func NewBankController(service primary.BankService) *BankController {
	return &BankController{service: service}
}

func (c *BankController) AddCreditCard(ctx *gin.Context) {
	var card entities.CreditCard
	if err := ctx.ShouldBindJSON(&card); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.service.AddCreditCard(card)
	if err != nil {
		log.Error().Err(err).Msg("Error adding credit card")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Credit card added successfully"})
}

func (c *BankController) GetAllCreditCards(ctx *gin.Context) {
	cards, err := c.service.GetAllCreditCards()
	if err != nil {
		log.Error().Err(err).Msg("Error getting credit cards")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	ctx.JSON(http.StatusOK, cards)
}

func (c *BankController) Pay(ctx *gin.Context) {
	var request entities.PaymentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.service.Pay(request)
	if err != nil {
		log.Error().Err(err).Msg("Error processing payment")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully"})
}

func (c *BankController) Refund(ctx *gin.Context) {
	var request entities.RefundRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Error binding JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.service.Refund(request)
	if err != nil {
		log.Error().Err(err).Msg("Error processing refund")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Refund processed successfully"})
}
