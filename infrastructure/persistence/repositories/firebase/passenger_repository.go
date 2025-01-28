package firebase

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"context"
	"fmt"

	"firebase.google.com/go/v4/db"
)

var PASSENGER_LOG_PREFIX string = "passenger_repository.go"

type PassengerRepositoryImpl struct {
	client *db.Client
}

func NewPassengerRepositoryImpl(client *db.Client) ports.PassengerRepository {
	return &PassengerRepositoryImpl{client: client}
}

func (r *PassengerRepositoryImpl) GetPassengerByID(passengerID string) (entities.Passenger, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying passenger by ID: %s", PASSENGER_LOG_PREFIX, passengerID))

	ctx := context.Background()
	ref := r.client.NewRef("passengers").Child(passengerID)

	var passenger entities.Passenger
	err := ref.Get(ctx, &passenger)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying passenger by ID: %v", PASSENGER_LOG_PREFIX, err))
		return entities.Passenger{}, err
	}

	return passenger, nil
}

func (r *PassengerRepositoryImpl) OnlineCheckInPassenger(pnr, surname string) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Checking in passenger with PNR: %s and surname: %s", PASSENGER_LOG_PREFIX, pnr, surname))

	ctx := context.Background()
	ref := r.client.NewRef("passengers")

	var passengers map[string]entities.Passenger
	err := ref.OrderByChild("pnr_no").EqualTo(pnr).Get(ctx, &passengers)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error checking in passenger: %v", PASSENGER_LOG_PREFIX, err))
		return err
	}

	for key, passenger := range passengers {
		if passenger.Surname == surname {
			passenger.CheckIn = true
			err := ref.Child(key).Set(ctx, passenger)
			if err != nil {
				middlewares.LogError(fmt.Sprintf("%s - Error updating check-in status: %v", PASSENGER_LOG_PREFIX, err))
				return err
			}
			break
		}
	}

	return nil
}
