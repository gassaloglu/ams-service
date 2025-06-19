package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"
	"errors"

	"github.com/sourcegraph/conc/iter"
)

type BankServiceImpl struct {
	repo secondary.BankRepository
}

func NewBankService(repo secondary.BankRepository) primary.BankService {
	return &BankServiceImpl{repo: repo}
}

func (s *BankServiceImpl) CreateCreditCard(request *entities.CreateCreditCardRequest) (*entities.CreditCard, error) {
	card := mapCreateCreditCardRequestToCreditCardEntity(request)
	return s.repo.CreateCreditCard(&card)
}

func (s *BankServiceImpl) CreateAllCreditCards(requests []entities.CreateCreditCardRequest) error {
	cards := iter.Map(requests, mapCreateCreditCardRequestToCreditCardEntity)
	return s.repo.CreateAllCreditCards(cards)
}

func (s *BankServiceImpl) Pay(request *entities.PaymentRequest) (*entities.Transaction, error) {
	card, err := s.repo.FindCreditCard(&request.CreditCard)

	if err != nil {
		return nil, err
	}

	if request.Amount <= 0 {
		return nil, errors.New("invalid amount")
	}

	if card.Balance < request.Amount {
		return nil, errors.New("insufficient balance")
	}

	card.Balance -= request.Amount

	_, err = s.repo.UpdateCreditCard(&card)

	if err != nil {
		return nil, err
	}

	potentiallyFraud := s.assessFraudPotential(request, &card)

	transaction, err := s.repo.CreateTransaction(&entities.Transaction{
		CreditCardID:     card.ID,
		Amount:           request.Amount,
		Type:             "credit",
		PotentiallyFraud: potentiallyFraud,
	})

	return transaction, err
}

func (s *BankServiceImpl) Refund(request *entities.RefundRequest) (*entities.Transaction, error) {
	transaction, err := s.repo.FindTransactionById(request.TransactionID)

	if err != nil {
		return nil, err
	}

	if transaction.Type != "credit" {
		return nil, errors.New("transaction is not refundable")
	}

	transaction, err = s.repo.CreateTransaction(&entities.Transaction{
		CreditCardID: transaction.CreditCardID,
		Amount:       transaction.Amount,
		Type:         "refund",
	})

	return transaction, err
}

func mapCreateCreditCardRequestToCreditCardEntity(request *entities.CreateCreditCardRequest) entities.CreditCard {
	return entities.CreditCard{
		CardNumber:        request.CardNumber,
		CardHolderName:    request.CardHolderName,
		CardHolderSurname: request.CardHolderSurname,
		ExpirationMonth:   request.ExpirationMonth,
		ExpirationYear:    request.ExpirationYear,
		CVV:               request.CVV,
		Balance:           100_000.0,
	}
}

// GetAllTransactions implements primary.BankService.
func (s *BankServiceImpl) GetAllTransactions(request *entities.GetAllTransactionsRequest) ([]entities.Transaction, error) {
	return s.repo.GetAllTransactions(request)
}

// assessFraudPotential determines if a transaction is potentially fraudulent.
func (s *BankServiceImpl) assessFraudPotential(request *entities.PaymentRequest, card *entities.CreditCard) bool {
	// TODO: Implement fraud assessment logic
	return false
}

func (s *BankServiceImpl) GetAllFraudulentActivities() ([]entities.FraudulentActivity, error) {
	return s.repo.GetAllFraudulentActivities()
}
