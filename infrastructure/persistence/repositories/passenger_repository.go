package repositories

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"fmt"
)

var LOG_PREFIX string = "passenger_repository.go"

/* ADAPTER - HANDLER */

type PassengerRepositoryImpl struct {
	// DB connection details
}

func NewPassengerRepositoryImpl( /*DB connection parameters*/ ) ports.PassengerRepository {
	return &PassengerRepositoryImpl{ /* init db*/ }
}

func (r *PassengerRepositoryImpl) GetPassengerByID(passengerID string) (entities.Passenger, error) {
	// Query the database to find the passenger by ID
	middlewares.LogInfo(fmt.Sprintf("%s - Querying passenger by ID: %s", LOG_PREFIX, passengerID))
	// Implement actual database query here
	return entities.Passenger{}, nil // Return passenger entity or error if not found
}

func (r *PassengerRepositoryImpl) OnlineCheckInPassenger(pnr, surname string) error {
	// Update the passenger's check-in status in the database
	middlewares.LogInfo(fmt.Sprintf("%s - Checking in passenger with PNR: %s and surname: %s", LOG_PREFIX, pnr, surname))
	// Implement actual database update here
	return nil
}
