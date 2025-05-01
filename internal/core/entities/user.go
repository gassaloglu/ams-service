package entities

import (
	"time"
)

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"size:50;not null"`
	Surname      string    `json:"surname" gorm:"size:50;not null"`
	Email        string    `json:"email" gorm:"unique;size:100;not null"`
	Phone        string    `json:"phone" gorm:"size:15"`
	Gender       string    `json:"gender" gorm:"type:gender_enum;not null"`
	BirthDate    time.Time `json:"birth_date" gorm:"not null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	Salt         string    `json:"-" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
