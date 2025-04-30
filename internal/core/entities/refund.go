package entities

import "time"

type Refund struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	RefundID  string    `json:"refund_id" gorm:"unique;not null;size:50"`
	PaymentID string    `json:"payment_id" gorm:"not null;size:50;foreignKey:id;references:payments;constraint:OnDelete:RESTRICT"`
	Amount    float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency  string    `json:"currency" gorm:"size:3;not null"`
	Reason    string    `json:"reason" gorm:"size:255"`
	Status    string    `json:"status" gorm:"type:status_enum;default:'active';not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
