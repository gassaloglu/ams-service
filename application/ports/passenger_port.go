package ports

import "ams-passenger-service/core/entities"

type PassengerRepository interface {
	GetPassengerByID(passengerID string) (entities.Passenger, error)
	OnlineCheckInPassenger(pnr string, surname string) error
}
