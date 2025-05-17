package secondary

import (
	"ams-service/internal/core/entities"
)

type PassengerRepository interface {
	GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error)
	GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error)
	OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error
	GetPassengersBySpecificFlight(request entities.GetPassengersBySpecificFlightRequest) ([]entities.Passenger, error)
	CreatePassenger(request *entities.Passenger) (*entities.Passenger, error)
	GetAllPassengers() ([]entities.Passenger, error)
	EmployeeCheckInPassenger(request entities.EmployeeCheckInRequest) (entities.Passenger, error)
	CancelPassenger(request entities.CancelPassengerRequest) error
}
