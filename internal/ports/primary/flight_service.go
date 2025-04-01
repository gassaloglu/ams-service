package primary

import (
	"ams-service/internal/core/entities"
)

type FlightService interface {
	GetSpecificFlight(request entities.GetSpecificFlightRequest, userId string) (entities.Flight, error)
	GetAllFlights() ([]entities.Flight, error)
	GetAllSpecificFlights(request entities.GetSpecificFlightsRequest) ([]entities.Flight, error)
	GetAllActiveFlights() ([]entities.Flight, error)
	CancelFlight(request entities.CancelFlightRequest) error
}
