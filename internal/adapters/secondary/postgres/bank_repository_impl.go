package postgres

import (
	"ams-service/internal/core/entities"

	"gorm.io/gorm"
)

type BankRepositoryImpl struct {
	db *gorm.DB
}

func NewBankRepositoryImpl(db *gorm.DB) *BankRepositoryImpl {
	db.AutoMigrate(&entities.CreditCard{}, &entities.Transaction{})
	return &BankRepositoryImpl{db: db}
}

func (r *BankRepositoryImpl) CreateCreditCard(card *entities.CreditCard) (*entities.CreditCard, error) {
	err := r.db.Create(card).Error
	return card, err
}

func (r *BankRepositoryImpl) CreateAllCreditCards(cards []entities.CreditCard) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(cards)
		return result.Error
	})
}

func (r *BankRepositoryImpl) FindCreditCard(info *entities.CreditCardInfo) (entities.CreditCard, error) {
	var card entities.CreditCard

	result := r.db.
		Where("card_number", info.CardNumber).
		Where("card_holder_name", info.CardHolderName).
		Where("card_holder_surname", info.CardHolderSurname).
		Where("expiration_month", info.ExpirationMonth).
		Where("expiration_year", info.ExpirationYear).
		Where("cvv", info.CVV).
		First(&card)

	return card, result.Error
}

func (r *BankRepositoryImpl) UpdateCreditCard(card *entities.CreditCard) (*entities.CreditCard, error) {
	clone := *card
	err := r.db.Save(&clone).Error
	return &clone, err
}

func (r *BankRepositoryImpl) CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	clone := *transaction
	result := r.db.Create(&clone)
	return &clone, result.Error
}

func (r *BankRepositoryImpl) FindTransactionById(id uint) (*entities.Transaction, error) {
	var transaction entities.Transaction
	err := r.db.First(&transaction, "id = ?", id).Error
	return &transaction, err
}
