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
