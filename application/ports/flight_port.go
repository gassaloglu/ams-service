package ports

import "ams-service/core/entities"

type FlightRepository interface {
	GetSpecificFlight(request entities.GetSpecificFlightRequest) (entities.Flight, error)
	GetAllFlights() ([]entities.Flight, error)
}
