package entities

import (
	"time"
)

type Flight struct {
	ID                    uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	FlightNumber          string    `json:"flight_number" gorm:"size:10;not null;"`
	DepartureAirport      string    `json:"departure_airport" gorm:"size:3;not null"`
	DestinationAirport    string    `json:"destination_airport" gorm:"size:3;not null"`
	DepartureDatetime     time.Time `json:"departure_datetime" gorm:"not null;"`
	ArrivalDatetime       time.Time `json:"arrival_datetime" gorm:"not null"`
	DepartureGateNumber   string    `json:"departure_gate_number" gorm:"size:5"`
	DestinationGateNumber string    `json:"destination_gate_number" gorm:"size:5"`
	PlaneRegistration     string    `json:"plane_registration" gorm:"size:10;not null"`
	Status                string    `json:"status" gorm:"type:flight_status_enum;not null"`
	Price                 float64   `json:"price" gorm:"type:decimal(10,2);not null"`
}
