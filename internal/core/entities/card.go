package entities

import (
	"time"
)

type CreditCard struct {
	ID                uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CardNumber        string    `json:"card_number" gorm:"size:16;not null"` // encrypted
	CardHolderName    string    `json:"card_holder_name" gorm:"size:100;not null"`
	CardHolderSurname string    `json:"card_holder_surname" gorm:"size:100;not null"`
	ExpirationMonth   int       `json:"expiration_month" gorm:"not null"` // 1-12
	ExpirationYear    int       `json:"expiration_year" gorm:"not null"`
	CVV               string    `json:"cvv" gorm:"size:4;not null"` // encrypted, for amex size: 4
	CardType          string    `json:"card_type" gorm:"type:card_type_enum;not null"`
	Amount            float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency          string    `json:"currency" gorm:"size:3;not null"`
	IssuerBank        string    `json:"issuer_bank" gorm:"size:100"`
	Status            string    `json:"status" gorm:"type:status_enum;default:'active';not null"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"` // Request creation timestamp
}
