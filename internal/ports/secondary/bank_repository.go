package secondary

import (
	"ams-service/internal/core/entities"
)

type BankRepository interface {
	AddCreditCard(card entities.CreditCard) error
	GetAllCreditCards() ([]entities.CreditCard, error)
	Pay(request entities.PaymentRequest) error
	Refund(request entities.RefundRequest) error
}
