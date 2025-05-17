package entities

import "time"

type Passenger struct {
	// Entity properties
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Ticket properties
	PnrNo         string `json:"pnr_no" gorm:"size:6;unique;not null"`
	CheckIn       bool   `json:"check_in" gorm:"not null"`
	FlightId      uint   `json:"flight_id" gorm:"not null;foreignKey:id;references:flights;constraint:OnDelete:RESTRICT"`
	TransactionId uint   `json:"transaction_id" gorm:"not null;foreignKey:id;references:transactions;constraint:OnDelete:RESTRICT"`
	Status        string `json:"status" gorm:"type:status_enum;not null"`

	// Passenger properties
	NationalId       string    `json:"national_id" gorm:"size:11;not null"`
	BaggageAllowance int       `json:"baggage_allowance" gorm:"not null"`
	ExtraBaggage     int       `json:"extra_baggage" gorm:"not null"`
	FareType         string    `json:"fare_type" gorm:"size:50;not null"`
	Meal             string    `json:"meal" gorm:"size:50;not null"`
	Name             string    `json:"name" gorm:"size:50;not null"`
	Surname          string    `json:"surname" gorm:"size:50;not null"`
	Email            string    `json:"email" gorm:"size:100;not null"`
	Phone            string    `json:"phone" gorm:"size:15;not null"`
	Gender           string    `json:"gender" gorm:"type:gender_enum;not null"`
	BirthDate        time.Time `json:"birth_date" gorm:"not null"`
	CipMember        bool      `json:"cip_member" gorm:"not null"`
	VipMember        bool      `json:"vip_member" gorm:"not null"`
	Disabled         bool      `json:"disabled" gorm:"not null"`
	Child            bool      `json:"child" gorm:"not null"`
	Seat             uint      `json:"seat" gorm:"default:null"`
}
