package services

import (
	"ams-service/core/entities"
	"ams-service/middlewares"
	"fmt"
)

var BANK_LOG_PREFIX string = "bank_service.go"

type BankRepository interface {
	AddCreditCard(card entities.CreditCard) error
	GetAllCreditCards() ([]entities.CreditCard, error)
	Pay(request entities.PaymentRequest) error
	Refund(request entities.RefundRequest) error
}

type BankService struct {
	repo BankRepository
}

func NewBankService(repo BankRepository) *BankService {
	return &BankService{repo: repo}
}

func (s *BankService) AddCreditCard(card entities.CreditCard) error {
	err := s.repo.AddCreditCard(card)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error adding credit card: %v", BANK_LOG_PREFIX, err))
		return err
	}
	return nil
}

func (s *BankService) GetAllCreditCards() ([]entities.CreditCard, error) {
	cards, err := s.repo.GetAllCreditCards()
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting credit cards: %v", BANK_LOG_PREFIX, err))
		return nil, err
	}
	return cards, nil
}

func (s *BankService) Pay(request entities.PaymentRequest) error {
	err := s.repo.Pay(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error processing payment: %v", BANK_LOG_PREFIX, err))
		return err
	}
	return nil
}

func (s *BankService) Refund(request entities.RefundRequest) error {
	err := s.repo.Refund(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error processing refund: %v", BANK_LOG_PREFIX, err))
		return err
	}
	return nil
}
