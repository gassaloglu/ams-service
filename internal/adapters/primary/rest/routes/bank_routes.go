package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterBankRoutes(app *fiber.App, bankController *controllers.BankController) {
	bankRoute := app.Group("/bank")
	bankRoute.Post("/card", bankController.AddCreditCard)
	bankRoute.Get("/cards", bankController.GetAllCreditCards)
	bankRoute.Post("/pay", bankController.Pay)
}
