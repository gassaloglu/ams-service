package controllers

import (
	"ams-service/core/entities"
	"ams-service/core/services"
	"ams-service/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

var BANK_LOG_PREFIX string = "bank_controller.go"

type BankController struct {
	service *services.BankService
}

func NewBankController(service *services.BankService) *BankController {
	return &BankController{service: service}
}

func (c *BankController) AddCreditCard(ctx *gin.Context) {
	var card entities.CreditCard
	if err := ctx.ShouldBindJSON(&card); err != nil {
		middlewares.LogError(BANK_LOG_PREFIX + " - Error binding JSON: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.service.AddCreditCard(card)
	if err != nil {
		middlewares.LogError(BANK_LOG_PREFIX + " - Error adding credit card: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Credit card added successfully"})
}

func (c *BankController) GetAllCreditCards(ctx *gin.Context) {
	cards, err := c.service.GetAllCreditCards()
	if err != nil {
		middlewares.LogError(BANK_LOG_PREFIX + " - Error getting credit cards: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	ctx.JSON(http.StatusOK, cards)
}

func (c *BankController) Pay(ctx *gin.Context) {
	var request entities.PaymentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		middlewares.LogError(BANK_LOG_PREFIX + " - Error binding JSON: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.service.Pay(request)
	if err != nil {
		middlewares.LogError(BANK_LOG_PREFIX + " - Error processing payment: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully"})
}

func (c *BankController) Refund(ctx *gin.Context) {
	var request entities.RefundRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		middlewares.LogError(BANK_LOG_PREFIX + " - Error binding JSON: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := c.service.Refund(request)
	if err != nil {
		middlewares.LogError(BANK_LOG_PREFIX + " - Error processing refund: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Refund processed successfully"})
}
