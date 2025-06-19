package secondary

import (
	"ams-service/internal/core/entities"
)

type BankRepository interface {
	CreateCreditCard(card *entities.CreditCard) (*entities.CreditCard, error)
	CreateAllCreditCards(requests []entities.CreditCard) error
	FindCreditCard(info *entities.CreditCardInfo) (entities.CreditCard, error)
	UpdateCreditCard(card *entities.CreditCard) (*entities.CreditCard, error)
	CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error)
	FindTransactionById(id uint) (*entities.Transaction, error)
	GetAllTransactions(request *entities.GetAllTransactionsRequest) ([]entities.Transaction, error)
	GetAllFraudulentActivities() ([]entities.FraudulentActivity, error)
}
