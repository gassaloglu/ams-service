package primary

import "ams-service/internal/core/entities"

type BankService interface {
	AddCreditCard(card entities.CreditCard) error
	GetAllCreditCards() ([]entities.CreditCard, error)
	Pay(request entities.PaymentRequest) error
	Refund(request entities.RefundRequest) error
}
