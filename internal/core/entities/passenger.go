package entities

import "time"

type Passenger struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Status           string    `json:"status" gorm:"type:status_enum;not null"`
	NationalId       string    `json:"national_id" gorm:"size:11;not null"`
	PnrNo            string    `json:"pnr_no" gorm:"size:6;unique;not null"`
	FlightId         int       `json:"flight_id" gorm:"not null;foreignKey:id;references:flights;constraint:OnDelete:RESTRICT"`
	PaymentId        int       `json:"payment_id" gorm:"not null;foreignKey:id;references:payments;constraint:OnDelete:RESTRICT"`
	BaggageAllowance int       `json:"baggage_allowance" gorm:"not null"`
	ExtraBaggage     int       `json:"extra_baggage" gorm:"not null"`
	FareType         string    `json:"fare_type" gorm:"size:50;not null"`
	Seat             int       `json:"seat" gorm:"default:null"`
	Meal             string    `json:"meal" gorm:"size:50;not null"`
	CheckIn          bool      `json:"check_in" gorm:"not null"`
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
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
