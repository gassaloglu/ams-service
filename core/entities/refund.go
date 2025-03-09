package entities

type RefundRequest struct {
	PaymentID uint    `json:"payment_id" gorm:"not null"`
	Amount    float64 `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency  string  `json:"currency" gorm:"size:3;not null"`
	Reason    string  `json:"reason" gorm:"size:255"`
	Status    string  `json:"status" gorm:"type:status_enum;default:'pending';not null"`
}
