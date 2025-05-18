package secondary

import (
	"ams-service/internal/core/entities"
)

type FlightRepository interface {
	FindById(id string) (*entities.Flight, error)
	FindByFlightNumber(flightNumber string) (*entities.Flight, error)
	FindAll(*entities.GetAllFlightsRequest) ([]entities.Flight, error)
	FindAllActive(*entities.GetAllFlightsRequest) ([]entities.Flight, error)
	Create(flight *entities.Flight) error
	CreateAll(flights []entities.Flight) error
	FindSeatsByFlightId(id string) ([]int, error)
}
