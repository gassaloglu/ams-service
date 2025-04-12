package postgres

import (
	"ams-service/internal/core/entities"
	"database/sql"
	"strconv"

	"github.com/rs/zerolog/log"
)

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
		log.Error().Err(err).Str("card_holder", card.CardHolderName).Msg("Error adding credit card")
		return err
	}
	return nil
}

func (r *BankRepositoryImpl) GetAllCreditCards() ([]entities.CreditCard, error) {
	query := `SELECT id, card_holder_name, card_holder_surname, expiration_month, expiration_year, card_type, amount, currency, issuer_bank, status, created_at FROM credit_cards`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Error().Err(err).Msg("Error getting credit cards")
		return nil, err
	}
	defer rows.Close()

	var cards []entities.CreditCard
	for rows.Next() {
		var card entities.CreditCard
		if err := rows.Scan(&card.ID, &card.CardHolderName, &card.CardHolderSurname, &card.ExpirationMonth, &card.ExpirationYear, &card.CardType, &card.Amount, &card.Currency, &card.IssuerBank, &card.Status, &card.CreatedAt); err != nil {
			log.Error().Err(err).Msg("Error scanning credit card")
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
		log.Error().Err(err).Str("transaction_id", request.TransactionID).Msg("Error processing payment")
		return err
	}
	return nil
}

func (r *BankRepositoryImpl) Refund(request entities.RefundRequest) error {
	query := `INSERT INTO refunds (payment_id, amount, currency, reason, status) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, request.PaymentID, request.Amount, request.Currency, request.Reason, request.Status)
	if err != nil {
		log.Error().Err(err).Str("payment_id", strconv.FormatUint(uint64(request.ID), 10)).Msg("Error processing refund")
		return err
	}
	return nil
}
