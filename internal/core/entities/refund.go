package entities

import "time"

type RefundRequest struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	RefundID  string    `json:"refund_id" gorm:"unique;not null;size:50"`
	PaymentID string    `json:"payment_id" gorm:"not null;size:50"` // Changed from uint to string to match VARCHAR(50)
	Amount    float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency  string    `json:"currency" gorm:"size:3;not null"`
	Reason    string    `json:"reason" gorm:"size:255"`
	Status    string    `json:"status" gorm:"type:status_enum;default:'active';not null"` // Updated default to 'active'
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
