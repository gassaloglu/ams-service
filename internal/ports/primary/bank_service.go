package primary

import "ams-service/internal/core/entities"

type BankService interface {
	CreateCreditCard(request *entities.CreateCreditCardRequest) (*entities.CreditCard, error)
	Pay(request *entities.PaymentRequest) (*entities.Transaction, error)
	Refund(request *entities.RefundRequest) (*entities.Transaction, error)
}
