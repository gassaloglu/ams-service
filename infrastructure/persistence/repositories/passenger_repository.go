package repositories

import (
	"ams-passenger-service/application/ports"
	"ams-passenger-service/core/entities"
)

/* ADAPTER - HANDLER */

type PassengerRepositoryImpl struct {
	// DB connection details
}

func NewPassengerRepositoryImpl( /*DB connection parameters*/ ) ports.PassengerRepository {
	return &PassengerRepositoryImpl{ /* init db*/ }
}

func (r *PassengerRepositoryImpl) GetPassengerByID(passengerID string) (entities.Passenger, error) {
	// Query the database to find the passenger by ID and surname
	return entities.Passenger{}, nil // Return passenger entity or error if not found
}

func (r *PassengerRepositoryImpl) OnlineCheckInPassenger(passengerID, surname string) error {
	// Update the passenger's check-in status in the database
	return nil
}
