package entities

import (
	"time"
)

type PaymentRequest struct {
	ID                uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Amount            float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	CardHolderName    string    `json:"card_holder_name" gorm:"size:100;not null"`
	CardHolderSurname string    `json:"card_holder_surname" gorm:"size:100;not null"`
	ExpirationMonth   int       `json:"expiration_month" gorm:"not null"`
	ExpirationYear    int       `json:"expiration_year" gorm:"not null"`
	Currency          string    `json:"currency" gorm:"size:3;not null"`
	IssuerBank        string    `json:"issuer_bank" gorm:"size:100"`
	Status            string    `json:"status" gorm:"type:status_enum;default:'active';not null"`
	TransactionID     string    `json:"transaction_id" gorm:"size:100;not null"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
}
