package entities

import "time"

type Transaction struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreditCardID     uint      `json:"credit_card_id" gorm:"not null"`
	Type             string    `json:"type" gorm:"type:transaction_type_enum;not null"`
	Amount           float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	PotentiallyFraud bool      `gorm:"not null;default:false" json:"potentially_fraud"`
}
