package entities

import "time"

type Plane struct {
	Registration string    `json:"registration" gorm:"primaryKey;size:10;not null"`
	Model        string    `json:"model" gorm:"size:50;not null"`
	Manufacturer string    `json:"manufacturer" gorm:"size:50;not null"`
	Capacity     int       `json:"capacity" gorm:"not null"`
	Status       string    `json:"status" gorm:"type:status_enum;default:'active';not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
