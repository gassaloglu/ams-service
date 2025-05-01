package entities

import (
	"time"
)

type Employee struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	NationalID   string    `json:"national_id" gorm:"type:varchar(50);unique;not null"`
	Name         string    `json:"name" gorm:"type:varchar(50);not null"`
	Surname      string    `json:"surname" gorm:"type:varchar(50);not null"`
	Email        string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Phone        string    `json:"phone" gorm:"type:varchar(15)"`
	Gender       string    `json:"gender" gorm:"type:gender_enum;not null"`
	BirthDate    time.Time `json:"birth_date" gorm:"type:timestamp;not null"`
	PasswordHash string    `json:"-" gorm:"type:text;not null"`
	Salt         string    `json:"-" gorm:"type:text;not null"`
	Title        string    `json:"title" gorm:"type:varchar(100);not null"`
	Role         string    `json:"role" gorm:"type:role_enum;not null"`
	Status       string    `json:"status" gorm:"type:status_enum;default:'active';not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
