package app

import (
	"ams-service/infrastructure/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBankRoutes(router *gin.Engine, bankController *controllers.BankController) {
	bankRoute := router.Group("/bank")
	{
		bankRoute.POST("/card/add", bankController.AddCreditCard)
		bankRoute.GET("/card/all", bankController.GetAllCreditCards)
		bankRoute.POST("/pay", bankController.Pay)
		bankRoute.POST("/refund", bankController.Refund)
	}
}
