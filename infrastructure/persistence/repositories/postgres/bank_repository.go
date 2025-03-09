package postgres

import (
	"ams-service/core/entities"
	"ams-service/middlewares"
	"database/sql"
	"fmt"
)

var BANK_LOG_PREFIX string = "bank_repository.go"

type BankRepositoryImpl struct {
	db *sql.DB
}

func NewBankRepositoryImpl(db *sql.DB) *BankRepositoryImpl {
	return &BankRepositoryImpl{db: db}
}

func (r *BankRepositoryImpl) AddCreditCard(card entities.CreditCard) error {
	query := `INSERT INTO credit_cards (card_number, card_holder_name, card_holder_surname, expiration_month, expiration_year, cvv, card_type, amount, currency, issuer_bank, status, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := r.db.Exec(query, card.CardNumber, card.CardHolderName, card.CardHolderSurname, card.ExpirationMonth, card.ExpirationYear, card.CVV, card.CardType, card.Amount, card.Currency, card.IssuerBank, card.Status, card.CreatedAt)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error adding credit card: %v", BANK_LOG_PREFIX, err))
		return err
	}
	return nil
}

func (r *BankRepositoryImpl) GetAllCreditCards() ([]entities.CreditCard, error) {
	query := `SELECT id, card_holder_name, card_holder_surname, expiration_month, expiration_year, card_type, amount, currency, issuer_bank, status, created_at FROM credit_cards`
	rows, err := r.db.Query(query)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting credit cards: %v", BANK_LOG_PREFIX, err))
		return nil, err
	}
	defer rows.Close()

	var cards []entities.CreditCard
	for rows.Next() {
		var card entities.CreditCard
		if err := rows.Scan(&card.ID, &card.CardHolderName, &card.CardHolderSurname, &card.ExpirationMonth, &card.ExpirationYear, &card.CardType, &card.Amount, &card.Currency, &card.IssuerBank, &card.Status, &card.CreatedAt); err != nil {
			middlewares.LogError(fmt.Sprintf("%s - Error scanning credit card: %v", BANK_LOG_PREFIX, err))
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (r *BankRepositoryImpl) Pay(request entities.PaymentRequest) error {
	query := `INSERT INTO payments (payment_request_id, status, transaction_id, amount, currency) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, request.ID, request.Status, request.TransactionID, request.Amount, request.Currency)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error processing payment: %v", BANK_LOG_PREFIX, err))
		return err
	}
	return nil
}

func (r *BankRepositoryImpl) Refund(request entities.RefundRequest) error {
	query := `INSERT INTO refunds (payment_id, amount, currency, reason, status) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, request.PaymentID, request.Amount, request.Currency, request.Reason, request.Status)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error processing refund: %v", BANK_LOG_PREFIX, err))
		return err
	}
	return nil
}
