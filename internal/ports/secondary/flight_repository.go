package secondary

import (
	"ams-service/internal/core/entities"
)

type FlightRepository interface {
	GetSpecificFlight(request entities.GetSpecificFlightRequest) (entities.Flight, error)
	GetAllFlights() ([]entities.Flight, error)
	GetAllFlightsDestinationDateFlights(request entities.GetAllFlightsDestinationDateRequest) ([]entities.Flight, error)
	GetAllActiveFlights() ([]entities.Flight, error)
	CancelFlight(request entities.CancelFlightRequest) error // New method
}
