package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterBankRoutes(app *fiber.App, bankController *controllers.BankController) {
	bankRoute := app.Group("/bank")
	bankRoute.Post("/card/add", bankController.AddCreditCard)
	bankRoute.Get("/card/all", bankController.GetAllCreditCards)
	bankRoute.Post("/pay", bankController.Pay)
	bankRoute.Post("/refund", bankController.Refund)
}
