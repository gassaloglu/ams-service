package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"

	"github.com/rs/zerolog/log"
)

type BankServiceImpl struct {
	repo secondary.BankRepository
}

func NewBankService(repo secondary.BankRepository) primary.BankService {
	return &BankServiceImpl{repo: repo}
}

func (s *BankServiceImpl) AddCreditCard(card entities.CreditCard) error {
	err := s.repo.AddCreditCard(card)
	if err != nil {
		log.Error().Err(err).Msg("Error adding credit card")
		return err
	}
	return nil
}

func (s *BankServiceImpl) GetAllCreditCards() ([]entities.CreditCard, error) {
	cards, err := s.repo.GetAllCreditCards()
	if err != nil {
		log.Error().Err(err).Msg("Error getting credit cards")
		return nil, err
	}
	return cards, nil
}

func (s *BankServiceImpl) Pay(request entities.PaymentRequest) error {
	err := s.repo.Pay(request)
	if err != nil {
		log.Error().Err(err).Msg("Error processing payment")
		return err
	}
	return nil
}

func (s *BankServiceImpl) Refund(request entities.RefundRequest) error {
	err := s.repo.Refund(request)
	if err != nil {
		log.Error().Err(err).Msg("Error processing refund")
		return err
	}
	return nil
}
