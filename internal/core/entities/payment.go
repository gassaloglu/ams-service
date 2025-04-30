package entities

import (
	"time"
)

type Payment struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID        string    `json:"user_id" gorm:"not null;size:50;index:idx_payments_user_id"`
	CardNumber    string    `json:"card_number" gorm:"size:16;not null"`
	Amount        float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency      string    `json:"currency" gorm:"size:3;not null"`
	PaymentMethod string    `json:"payment_method" gorm:"size:50;not null"`
	Status        string    `json:"status" gorm:"type:status_enum;default:'active';not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
