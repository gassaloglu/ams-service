package routes

import (
	"ams-service/internal/adapters/primary/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterBankRoutes(app *fiber.App, bankController *controllers.BankController) {
	app.Post("/creditcards", bankController.CreateCreditCard)
}
