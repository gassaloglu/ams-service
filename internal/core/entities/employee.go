package entities

import (
	"time"
)

type Employee struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	EmployeeID       string    `json:"employee_id" gorm:"type:varchar(50);unique;not null"`
	Name             string    `json:"name" gorm:"type:varchar(50);not null"`
	Surname          string    `json:"surname" gorm:"type:varchar(50);not null"`
	Email            string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Phone            string    `json:"phone" gorm:"type:varchar(15)"`
	Address          string    `json:"address" gorm:"type:varchar(255)"`
	Gender           string    `json:"gender" gorm:"type:gender_enum;not null"`
	BirthDate        time.Time `json:"birth_date" gorm:"type:timestamp;not null"`
	HireDate         time.Time `json:"hire_date" gorm:"type:timestamp;not null"`
	Position         string    `json:"position" gorm:"type:varchar(100);not null"`
	Role             string    `json:"role" gorm:"type:role_enum;not null"`
	Salary           float64   `json:"salary" gorm:"type:decimal(10,2);not null"`
	Status           string    `json:"status" gorm:"type:status_enum;default:'active';not null"`
	ManagerID        *uint     `json:"manager_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	EmergencyContact string    `json:"emergency_contact" gorm:"type:varchar(100)"`
	EmergencyPhone   string    `json:"emergency_phone" gorm:"type:varchar(15)"`
	ProfileImageURL  string    `json:"profile_image_url" gorm:"type:varchar(255)"`
	CreatedAt        time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	PasswordHash     string    `json:"password_hash" gorm:"type:text;not null"`
	Salt             string    `json:"-" gorm:"type:text;not null"`
}
