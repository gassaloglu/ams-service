package secondary

import (
	"ams-service/internal/core/entities"
)

type FlightRepository interface {
	FindById(id string) (*entities.Flight, error)
	FindAll() ([]entities.Flight, error)
	FindAllActive() ([]entities.Flight, error)
	Create(flight *entities.Flight) error
	CreateAll(flights []entities.Flight) error
	FindSeatsByFlightId(id string) ([]int, error)
}
