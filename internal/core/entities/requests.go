package entities

import (
	"time"
)

/* Passenger */
type GetPassengerByIdRequest struct {
	NationalId string `query:"national_id" binding:"required,len=11,numeric"`
}

type GetPassengerByPnrRequest struct {
	PNR     string `query:"pnr" binding:"required,len=6,alphanum"`
	Surname string `query:"surname" binding:"required,alpha,min=2,max=50"`
}

type OnlineCheckInRequest struct {
	PNR     string `json:"pnr" binding:"len=6,alphanum"`
	Surname string `json:"surname" binding:"alpha,min=2,max=50"`
}

type EmployeeCheckInRequest struct {
	NationalId         string `json:"national_id" binding:"required,len=11,numeric"`
	DestinationAirport string `json:"destination_airport" binding:"required,len=3,alpha"`
}

type GetPassengersBySpecificFlightRequest struct {
	FlightNumber      string `query:"flight_number" binding:"required,len=6,alphanum"`
	DepartureDateTime string `query:"departure_datetime" binding:"required,datetime=2006-01-02"`
}

type CreatePassengerRequest struct {
	Passenger Passenger
}

/* Employee */
type RegisterEmployeeRequest []Employee

type LoginEmployeeRequest struct {
	EmployeeID string `json:"employee_id"`
	Password   string `json:"password"`
}

type GetEmployeeByIdRequest struct {
	EmployeeID string `json:"employee_id" binding:"required"`
}

/* Plane */
type AddPlaneRequest []Plane

type SetPlaneStatusRequest struct {
	PlaneRegistration string `params:"registration" binding:"required"`
	Status            string `json:"status" binding:"required"`
}

type GetAllPlanesRequest struct {
	Model        string `query:"model"`
	Manufacturer string `query:"manufacturer"`
	Capacity     int    `query:"capacity"`
	Status       string `query:"status"`
}

type GetPlaneByRegistrationRequest struct {
	Registration string `params:"registration" binding:"required"`
}

type GetPlaneByFlightNumberRequest struct {
	FlightNumber string `json:"flight_number" binding:"required"`
}

type GetPlaneByLocationRequest struct {
	Location string `json:"location" binding:"required"`
}

/* Flight */
type GetSpecificFlightRequest struct {
	FlightNumber      string `query:"flight_number"`
	DepartureDateTime string `query:"departure_datetime"`
}

type CancelFlightRequest struct {
	FlightNumber string `json:"flight_number" binding:"required"`
	FlightDate   string `json:"flight_date" binding:"required,datetime=2006-01-02"`
}

type AddFlightRequest []Flight

type GetAllFlightsDestinationDateRequest struct {
	DepartureAirport   string `query:"departure_airport" binding:"required,len=3,alpha"`
	DestinationAirport string `query:"destination_airport" binding:"required,len=3,alpha"`
	DepartureDateTime  string `query:"departure_datetime" binding:"required,datetime=2006-01-02"`
}

/* User */
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

/* Payment */
type PaymentRequest struct {
	ID                uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Amount            float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	CardNumber        string    `json:"card_number" gorm:"size:16;not null"` // encrypted
	CardHolderName    string    `json:"card_holder_name" gorm:"size:100;not null"`
	CardHolderSurname string    `json:"card_holder_surname" gorm:"size:100;not null"`
	ExpirationMonth   int       `json:"expiration_month" gorm:"not null"` // 1-12
	ExpirationYear    int       `json:"expiration_year" gorm:"not null"`
	CVV               string    `json:"cvv" gorm:"size:4;not null"` // encrypted, for amex size: 4
	Currency          string    `json:"currency" gorm:"size:3;not null"`
	IssuerBank        string    `json:"issuer_bank" gorm:"size:100"`
	Status            string    `json:"status" gorm:"type:status_enum;default:'active';not null"`
	TransactionID     string    `json:"transaction_id" gorm:"size:100;not null"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
}
