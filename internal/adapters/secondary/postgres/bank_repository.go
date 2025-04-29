package postgres

import (
	"ams-service/internal/core/entities"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type BankRepositoryImpl struct {
	db *gorm.DB
}

func NewBankRepositoryImpl(db *gorm.DB) *BankRepositoryImpl {
	db.AutoMigrate(&entities.CreditCard{})
	return &BankRepositoryImpl{db: db}
}

func (r *BankRepositoryImpl) AddCreditCard(card entities.CreditCard) error {
	result := r.db.Create(&card)
	return result.Error
}

func (r *BankRepositoryImpl) GetAllCreditCards() ([]entities.CreditCard, error) {
	var cards []entities.CreditCard
	result := r.db.Find(&cards)
	return cards, result.Error
}

func (r *BankRepositoryImpl) Pay(request entities.PaymentRequest) error {
	log.Fatal().Msg("todo: pay")
	return nil
}

func (r *BankRepositoryImpl) Refund(request entities.RefundRequest) error {
	log.Fatal().Msg("todo: refund")
	return nil
}
