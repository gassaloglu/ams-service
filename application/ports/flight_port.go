package ports

import "ams-service/core/entities"

type FlightRepository interface {
	GetSpecificFlight(request entities.GetSpecificFlightRequest) (entities.Flight, error)
	GetAllFlights() ([]entities.Flight, error)
	GetAllSpecificFlights(request entities.GetSpecificFlightsRequest) ([]entities.Flight, error)
	GetAllActiveFlights() ([]entities.Flight, error)
	CancelFlight(request entities.CancelFlightRequest) error // New method
}
