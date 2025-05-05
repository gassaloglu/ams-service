package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"
	"errors"
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

	transaction, err := s.repo.CreateTransaction(&entities.Transaction{
		CreditCardID: card.ID,
		Amount:       request.Amount,
		Type:         "credit",
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
		Balance:           0,
	}
}
