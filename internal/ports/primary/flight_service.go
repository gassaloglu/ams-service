package primary

import (
	"ams-service/internal/core/entities"
)

type FlightService interface {
	FindAll() ([]entities.Flight, error)
	FindAllActive() ([]entities.Flight, error)
	FindById(request *entities.GetFlightByIdRequest) (*entities.Flight, error)
	Create(request *entities.CreateFlightRequest) error
	CreateAll(requests []entities.CreateFlightRequest) error
	FindSeatsByFlightId(request *entities.GetSeatsByFlightIdRequest) ([]int, error)
}
