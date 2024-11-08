package ports

import "ams-service/core/entities"

type PassengerRepository interface {
	GetPassengerByID(passengerID string) (entities.Passenger, error)
	OnlineCheckInPassenger(pnr string, surname string) error
}
