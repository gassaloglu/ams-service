package entities

import (
	"time"
)

/* Util */
type Comparable[T comparable] struct {
	GreaterThan          *T `query:"gt"`
	LessThan             *T `query:"lt"`
	EqualTo              *T `query:"eq"`
	GreaterThanOrEqualTo *T `query:"gte"`
	LessThanOrEqualTo    *T `query:"lte"`
	NotEqaualTo          *T `query:"neq"`
}

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

type PassengerInfo struct {
	FlightID   uint      `json:"flight_id" binding:"required,alphanum,max=10"`
	FareType   string    `json:"fare_type" binding:"required,max=50"`
	NationalID string    `json:"national_id" binding:"required,len=11,numeric"`
	Name       string    `json:"name" binding:"required,alpha,min=2,max=50"`
	Surname    string    `json:"surname" binding:"required,alpha,min=2,max=50"`
	Email      string    `json:"email" binding:"required,email,max=100"`
	Phone      string    `json:"phone" binding:"required,e164,max=15"`
	Gender     string    `json:"gender" binding:"required,oneof=male female"`
	Disabled   bool      `json:"disabled" binding:"required,oneof=0 1"`
	Seat       uint      `json:"seat" binding:"required,numeric"`
	BirthDate  time.Time `json:"birth_date" binding:"required,datetime=2006-01-02"`
	Child      bool      `json:"child" binding:"required,oneof=0 1"`
}

type CreatePassengerRequest struct {
	Passenger  PassengerInfo  `json:"passenger" binding:"required"`
	CreditCard CreditCardInfo `json:"credit_card" binding:"required"`
}

type CancelPassengerRequest struct {
	PassengerID uint `json:"passenger_id" binding:"required"`
}

/* Employee */
type RegisterEmployeeRequest struct {
	NationalID string    `json:"national_id"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Gender     string    `json:"gender"`
	BirthDate  time.Time `json:"birth_date"`
	Password   string    `json:"password"`
	Title      string    `json:"title"`
	Role       string    `json:"role"`
}

type LoginEmployeeRequest struct {
	NationalID string `json:"national_id"`
	Password   string `json:"password"`
}

/* Plane */
type CreatePlaneRequest struct {
	Registration string `json:"registration"`
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	Capacity     int    `json:"capacity"`
	Status       string `json:"status"`
}

type GetAllPlanesRequest struct {
	Model        string `query:"model"`
	Manufacturer string `query:"manufacturer"`
	Capacity     int    `query:"capacity"`
	Status       string `query:"status"`
}

/* Flight */
type GetAllFlightsRequest struct {
	ID                 []uint                 `query:"id"`
	FlightNumber       []string               `query:"flight_number"`
	DepartureAirport   []string               `query:"departure_airport"`
	DestinationAirport []string               `query:"destination_airport"`
	DepartureDatetime  *Comparable[time.Time] `query:"departure_datetime"`
	Status             []string               `query:"status"`
	Price              *Comparable[float64]   `query:"price"`
}

type GetFlightByIdRequest struct {
	ID string `params:"id" binding:"required"`
}

type CreateFlightRequest struct {
	FlightNumber          string    `json:"flight_number"`
	DepartureAirport      string    `json:"departure_airport"`
	DestinationAirport    string    `json:"destination_airport"`
	DepartureDatetime     time.Time `json:"departure_datetime"`
	ArrivalDatetime       time.Time `json:"arrival_datetime"`
	DepartureGateNumber   string    `json:"departure_gate_number"`
	DestinationGateNumber string    `json:"destination_gate_number"`
	PlaneRegistration     string    `json:"plane_registration"`
	Price                 float64   `json:"price"`
}

type GetSeatsByFlightIdRequest struct {
	ID string `params:"id" binding:"required"`
}

/* User */
type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserRequest struct {
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json:"surname" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Phone     string    `json:"phone" binding:"required"`
	Gender    string    `json:"gender" binding:"required"`
	BirthDate time.Time `json:"birth_date" binding:"required,datetime=2006-01-02"`
}

/* Bank */

type CreditCardInfo struct {
	CardNumber        string `json:"card_number"`
	CardHolderName    string `json:"card_holder_name"`
	CardHolderSurname string `json:"card_holder_surname"`
	ExpirationMonth   int    `json:"expiration_month"`
	ExpirationYear    int    `json:"expiration_year"`
	CVV               string `json:"cvv"`
}

type CreateCreditCardRequest = CreditCardInfo

type PaymentRequest struct {
	CreditCard CreditCardInfo `json:"credit_card"`
	Amount     float64        `json:"amount"`
}

type RefundRequest struct {
	TransactionID string `json:"transaction_id" binding:"required"`
}
