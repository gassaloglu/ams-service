package entities

type GetPassengerByIdRequest struct {
	NationalId string `form:"national_id" binding:"required,len=11,numeric"`
}

type GetPassengerByPnrRequest struct {
	PNR     string `form:"pnr" binding:"required,len=6,alphanum"`
	Surname string `form:"surname" binding:"required,alpha,min=2,max=50"`
}

type OnlineCheckInRequest struct {
	PNR     string `json:"pnr" binding:"len=6,alphanum"`
	Surname string `json:"surname" binding:"alpha,min=2,max=50"`
}

// Will be updated(needs work)
type EmployeeCheckInRequest struct {
	PNR     string `json:"pnr" binding:"len=6,alphanum"`
	Surname string `json:"surname" binding:"alpha,min=2,max=50"`
}

type GetPassengersBySpecificFlightRequest struct {
	FlightNumber      string `form:"flight_number" binding:"required,len=6,alphanum"`
	DepartureDateTime string `form:"departure_datetime" binding:"required,datetime=2006-01-02"`
}

type CreatePassengerRequest struct {
	NationalId       string `json:"national_id" binding:"required,len=11,numeric"`
	PnrNo            string `json:"pnr_no" binding:"required,len=6,alphanum"`
	BaggageAllowance int    `json:"baggage_allowance" binding:"required"`
	BaggageId        string `json:"baggage_id" binding:"required"`
	FareType         string `json:"fare_type" binding:"required"`
	Seat             string `json:"seat" binding:"required"`
	Meal             string `json:"meal" binding:"required"`
	ExtraBaggage     int    `json:"extra_baggage" binding:"required"`
	CheckIn          bool   `json:"check_in"`
	Name             string `json:"name" binding:"required"`
	Surname          string `json:"surname" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	Phone            string `json:"phone" binding:"required,len=10,numeric"`
	Gender           string `json:"gender" binding:"required"`
	BirthDate        string `json:"birth_date"`
	CipMember        bool   `json:"cip_member"`
	VipMember        bool   `json:"vip_member"`
	Disabled         bool   `json:"disabled"`
	Child            bool   `json:"child"`
}

type GetEmployeeByIdRequest struct {
	ID uint `json:"id" binding:"required"`
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
	PlaneRegistration string `form:"registration_code" binding:"required"`
}

type GetPlaneByFlightNumberRequest struct {
	FlightNumber string `json:"flight_number" binding:"required"`
}

type GetPlaneByLocationRequest struct {
	Location string `json:"location" binding:"required"`
}

type GetSpecificFlightRequest struct {
	FlightNumber      string `form:"flight_number"`
	DepartureDateTime string `form:"departure_datetime"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetSpecificFlightsRequest struct {
	DepartureAirport   string `form:"departure_airport" binding:"required,len=3,alpha"`
	DestinationAirport string `form:"destination_airport" binding:"required,len=3,alpha"`
	DepartureDateTime  string `form:"departure_datetime" binding:"required,datetime=2006-01-02"`
}

type CancelFlightRequest struct {
	FlightNumber string `json:"flight_number" binding:"required"`
	FlightDate   string `json:"flight_date" binding:"required,datetime=2006-01-02"`
}
