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

func (r *BankRepositoryImpl) GetAllTransactions(request *entities.GetAllTransactionsRequest) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	db := r.db.Model(&entities.Transaction{})
	if request != nil {
		if len(request.ID) > 0 {
			db = db.Where("id IN ?", request.ID)
		}
		if len(request.CreditCardID) > 0 {
			db = db.Where("credit_card_id IN ?", request.CreditCardID)
		}
		if len(request.Type) > 0 {
			db = db.Where("type IN ?", request.Type)
		}
		if request.Amount != nil {
			if request.Amount.GreaterThan != nil {
				db = db.Where("amount > ?", *request.Amount.GreaterThan)
			}
			if request.Amount.LessThan != nil {
				db = db.Where("amount < ?", *request.Amount.LessThan)
			}
			if request.Amount.EqualTo != nil {
				db = db.Where("amount = ?", *request.Amount.EqualTo)
			}
			if request.Amount.GreaterThanOrEqualTo != nil {
				db = db.Where("amount >= ?", *request.Amount.GreaterThanOrEqualTo)
			}
			if request.Amount.LessThanOrEqualTo != nil {
				db = db.Where("amount <= ?", *request.Amount.LessThanOrEqualTo)
			}
			if request.Amount.NotEqaualTo != nil {
				db = db.Where("amount != ?", *request.Amount.NotEqaualTo)
			}
		}
		if request.PotentiallyFraud != nil {
			db = db.Where("potentially_fraud = ?", *request.PotentiallyFraud)
		}
	}
	err := db.Find(&transactions).Error
	return transactions, err
}

func (r *BankRepositoryImpl) GetAllFraudulentActivities() ([]entities.FraudulentActivity, error) {
	var results []entities.FraudulentActivity

	// Single GORM query with joins
	err := r.db.Table("transactions").
		Select("transactions.*, passengers.*, users.*").
		Joins("JOIN passengers ON passengers.transaction_id = transactions.id").
		Joins("JOIN users ON users.email = passengers.email").
		Where("transactions.potentially_fraud = ?", true).
		Scan(&results).Error

	if err == gorm.ErrRecordNotFound {
		return []entities.FraudulentActivity{}, nil
	}

	return results, err
}
