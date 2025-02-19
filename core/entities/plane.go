package entities

type Plane struct {
	PlaneRegistration    string `json:"plane_registration" binding:"required,alphanum"`
	PlaneType            string `json:"plane_type" binding:"required,alpha"`
	Location             string `json:"location" binding:"required,oneof=On-Flight Hangar Maintenance"`
	TotalPassengers      int    `json:"total_passengers" binding:"required"`
	MaxPassengers        int    `json:"max_passengers" binding:"required"`
	EconomyPassengers    int    `json:"economy_passengers" binding:"required"`
	BusinessPassengers   int    `json:"business_passengers" binding:"required,"`
	FirstClassPassengers int    `json:"first_class_passengers" binding:"required"`
	FlightNumber         string `json:"flight_number" binding:"required,alphanum"`
	IsAvailable          bool   `json:"is_available" binding:"required"`
}

type GetPlaneByRegistrationRequest struct {
	PlaneRegistration string `json:"plane_registration" binding:"required,alphanum"`
}

type GetPlaneByFlightNumberRequest struct {
	FlightNumber string `json:"flight_number" binding:"required,alphanum"`
}

type GetPlaneByLocationRequest struct {
	Location string `json:"location" binding:"required,oneof=On-Flight Hangar Maintenance"`
}

type AddPlaneRequest struct {
	Plane Plane
}

type SetPlaneStatusRequest struct {
	IsAvailable bool `json:"is_available" binding:"required"`
}
