package ports

import "ams-service/core/entities"

type PassengerRepository interface {
	GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error)
	GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error)
	OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error
	GetPassengersBySpecificFlight(request entities.GetPassengersBySpecificFlightRequest) ([]entities.Passenger, error)
	CreatePassenger(request entities.CreatePassengerRequest) error
	GetAllPassengers() ([]entities.Passenger, error)
}
