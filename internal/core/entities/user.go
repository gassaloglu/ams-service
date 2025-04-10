package entities

import (
	"time"
)

type User struct {
	ID                 uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name               string    `json:"name" gorm:"size:50;not null"`
	Surname            string    `json:"surname" gorm:"size:50;not null"`
	Username           string    `json:"username" gorm:"unique;size:50;not null"`
	Email              string    `json:"email" gorm:"unique;size:100;not null"`
	PasswordHash       string    `json:"password_hash" gorm:"not null"`
	Salt               string    `json:"-" gorm:"not null"`
	Phone              string    `json:"phone" gorm:"size:15"`
	Gender             string    `json:"gender" gorm:"type:gender_enum;not null"`
	BirthDate          time.Time `json:"birth_date" gorm:"not null"`
	LastLogin          time.Time `json:"last_login"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	LastPasswordChange time.Time `json:"last_password_change"`
}
