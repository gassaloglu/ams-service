package entities

type Flight struct {
	FlightNumber          string  `json:"flight_number" binding:"required,alphanum,max=6"`
	DepartureAirport      string  `json:"departure_airport" binding:"required,alpha,len=3"`
	DestinationAirport    string  `json:"destination_airport" binding:"required,alpha,len=3,nefield=DepartureAirport"`
	DepartureDateTime     string  `json:"departure_datetime" binding:"required,datetime=2006-01-02T15:04"`
	ArrivalDateTime       string  `json:"arrival_datetime" binding:"required,datetime=2006-01-02T15:04,gtfield=DepartureDateTime"`
	DepartureGateNumber   string  `json:"departure_gate_number" binding:"required,alphanum"`
	DestinationGateNumber string  `json:"destination_gate_number" binding:"required,alphanum,nefield=DepartureGateNumber"`
	PlaneRegistration     string  `json:"plane_registration" binding:"required,alphanum,len=6"`
	Status                string  `json:"status" binding:"required,oneof=scheduled delayed canceled departed arrived"`
	Price                 float32 `json:"price" binding:"required,gte=0"`
}

type GetSpecificFlightRequest struct {
	FlightNumber      string `form:"flightNumber" binding:"required"`
	DepartureDateTime string `form:"departureDateTime" binding:"required"`
}
