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

func (r *PassengerRepositoryImpl) GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying passenger by ID: %s", PASSENGER_LOG_PREFIX, request.NationalId))

	ctx := context.Background()
	ref := r.client.NewRef("passengers").Child(request.NationalId)

	var passenger entities.Passenger
	err := ref.Get(ctx, &passenger)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying passenger by ID: %v", PASSENGER_LOG_PREFIX, err))
		return entities.Passenger{}, err
	}

	return passenger, nil
}

func (r *PassengerRepositoryImpl) GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying passenger by PNR: %s and surname: %s", PASSENGER_LOG_PREFIX, request.PNR, request.Surname))

	ctx := context.Background()
	ref := r.client.NewRef("passengers").Child(request.PNR)

	var passenger entities.Passenger
	err := ref.Get(ctx, &passenger)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying passenger by PNR: %v", PASSENGER_LOG_PREFIX, err))
		return entities.Passenger{}, err
	}

	return passenger, nil
}

func (r *PassengerRepositoryImpl) OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Checking in passenger with PNR: %s and surname: %s", PASSENGER_LOG_PREFIX, request.PNR, request.Surname))

	ctx := context.Background()
	ref := r.client.NewRef("passengers")

	var passengers map[string]entities.Passenger
	err := ref.OrderByChild("pnr_no").EqualTo(request.PNR).Get(ctx, &passengers)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error checking in passenger: %v", PASSENGER_LOG_PREFIX, err))
		return err
	}

	for key, passenger := range passengers {
		if passenger.Surname == request.Surname {
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
