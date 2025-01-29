package entities

import (
	"time"
)

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"size:50;not null"`
	Surname      string    `json:"surname" gorm:"size:50;not null"`
	Username     string    `json:"username" gorm:"unique;size:50;not null"`
	Email        string    `json:"email" gorm:"unique;size:100;not null"`
	PasswordHash string    `json:"password_hash" gorm:"not null"`
	Salt         string    `json:"-" gorm:"not null"`
	Phone        string    `json:"phone" gorm:"size:15"`
	Gender       string    `json:"gender" gorm:"type:enum('male', 'female', 'other');not null"`
	BirthDate    time.Time `json:"birth_date" gorm:"not null"`
	Role         string    `json:"role" gorm:"type:enum('employee', 'user');default:'user';not null"`
	//IsActive           bool      `json:"is_active" gorm:"default:true"`
	//IsVerified         bool      `json:"is_verified" gorm:"default:false"`
	LastLogin time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	//PasswordResetToken string    `json:"-" gorm:"size:100"` // Helps manage password reset functionality with a token and expiry time.
	//ResetTokenExpiry   time.Time `json:"-"`
	//ProfileImageURL    string    `json:"profile_image_url" gorm:"size:255"`
	//Address            string    `json:"address" gorm:"size:255"`
	//LoginAttempts      int       `json:"login_attempts" gorm:"default:0"`
	LastPasswordChange time.Time `json:"last_password_change"`
}
