package entities

import (
	"time"
)

type CreditCard struct {
	ID                uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CardNumber        string    `json:"-" gorm:"size:16;not null"` // encrypted
	CardHolderName    string    `json:"card_holder_name" gorm:"size:100;not null"`
	CardHolderSurname string    `json:"card_holder_surname" gorm:"size:100;not null"`
	ExpirationMonth   int       `json:"expiration_month" gorm:"not null"` // 1-12
	ExpirationYear    int       `json:"expiration_year" gorm:"not null"`
	CVV               string    `json:"-" gorm:"size:4;not null"` // encrypted, for amex size: 4
	CardType          string    `json:"card_type" gorm:"type:enum('visa', 'mastercard', 'amex');not null"`
	Amount            float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency          string    `json:"currency" gorm:"size:3;not null"`
	IssuerBank        string    `json:"issuer_bank" gorm:"size:100"`
	Status            string    `json:"status" gorm:"type:enum('active', 'inactive', 'expired');default:'active';not null"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"` // Request creation timestamp
}

type PaymentRequest struct {
	ID                uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Amount            float64 `json:"amount" gorm:"type:decimal(10,2);not null"`
	CardHolderName    string  `json:"card_holder_name" gorm:"size:100;not null"`
	CardHolderSurname string  `json:"card_holder_surname" gorm:"size:100;not null"`
	ExpirationMonth   int     `json:"expiration_month" gorm:"not null"`
	ExpirationYear    int     `json:"expiration_year" gorm:"not null"`
	Status            string  `json:"status" gorm:"type:enum('approved', 'declined', 'failed');not null"`
	TransactionID     string  `json:"transaction_id" gorm:"unique;not null"`
	Currency          string  `json:"currency" gorm:"size:3;not null"`
}

type RefundRequest struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	PaymentID   uint      `json:"payment_id" gorm:"not null"` // ID of the associated PaymentResponse
	Amount      float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency    string    `json:"currency" gorm:"size:3;not null"`
	Reason      string    `json:"reason" gorm:"size:255"`
	RequestedAt time.Time `json:"requested_at" gorm:"autoCreateTime"`
	Status      string    `json:"status" gorm:"type:enum('pending', 'approved', 'rejected');default:'pending';not null"`
}

type RefundResponse struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	RefundRequestID uint      `json:"refund_request_id" gorm:"not null"` // ID of the associated RefundRequest
	Status          string    `json:"status" gorm:"type:enum('completed', 'failed');not null"`
	TransactionID   string    `json:"transaction_id" gorm:"unique;not null"`
	Amount          float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency        string    `json:"currency" gorm:"size:3;not null"`
	Message         string    `json:"message" gorm:"size:255"`
	ProcessedAt     time.Time `json:"processed_at" gorm:"autoCreateTime"`
}
