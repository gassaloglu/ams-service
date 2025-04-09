package entities

import (
	"time"
)

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
	NationalId       string    `json:"national_id" binding:"required,len=11,numeric"`
	PnrNo            string    `json:"pnr_no" binding:"required,len=6,alphanum"`
	FlightId         int       `json:"flight_id" binding:"required"`
	PaymentId        int       `json:"payment_id" binding:"required"`
	BaggageAllowance int       `json:"baggage_allowance" binding:"required"`
	BaggageId        string    `json:"baggage_id" binding:"required"`
	FareType         string    `json:"fare_type" binding:"required"`
	Seat             int       `json:"seat"`
	Meal             string    `json:"meal" binding:"required"`
	ExtraBaggage     int       `json:"extra_baggage" binding:"required"`
	CheckIn          bool      `json:"check_in"`
	Name             string    `json:"name" binding:"required"`
	Surname          string    `json:"surname" binding:"required"`
	Email            string    `json:"email" binding:"required,email"`
	Phone            string    `json:"phone" binding:"required,len=15,numeric"`
	Gender           string    `json:"gender" binding:"required"`
	BirthDate        time.Time `json:"birth_date"`
	CipMember        bool      `json:"cip_member"`
	VipMember        bool      `json:"vip_member"`
	Disabled         bool      `json:"disabled"`
	Child            bool      `json:"child"`
}

type RegisterEmployeeRequest struct {
	Employee Employee
}

type AddPlaneRequest struct {
	Plane Plane
}

type SetPlaneStatusRequest struct {
	PlaneRegistration string `json:"plane_registration" binding:"required"`
	IsAvailable       bool   `json:"is_available" binding:"required"`
}

type GetPlaneByRegistrationRequest struct {
	PlaneRegistration string `query:"registration_code" binding:"required"`
}

type GetPlaneByFlightNumberRequest struct {
	FlightNumber string `json:"flight_number" binding:"required"`
}

type GetPlaneByLocationRequest struct {
	Location string `json:"location" binding:"required"`
}

type GetSpecificFlightRequest struct {
	FlightNumber      string `query:"flight_number"`
	DepartureDateTime string `query:"departure_datetime"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginEmployeeRequest struct {
	EmployeeID string `json:"employee_id"`
	Password   string `json:"password"`
}

type GetAllFlightsDestinationDateRequest struct {
	DepartureAirport   string `query:"departure_airport" binding:"required,len=3,alpha"`
	DestinationAirport string `query:"destination_airport" binding:"required,len=3,alpha"`
	DepartureDateTime  string `query:"departure_datetime" binding:"required,datetime=2006-01-02"`
}

type CancelFlightRequest struct {
	FlightNumber string `json:"flight_number" binding:"required"`
	FlightDate   string `json:"flight_date" binding:"required,datetime=2006-01-02"`
}

type GetEmployeeByIdRequest struct {
	EmployeeID string `json:"employee_id" binding:"required"`
}

type AddFlightRequest struct {
	Flight Flight `json:"flight" binding:"required"`
}
