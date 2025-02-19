package firebase

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"context"
	"fmt"

	"firebase.google.com/go/v4/db"
)

var PLNAE_LOG_PREFIX string = "plane_repository.go"

type PlaneRepositoryImpl struct {
	client *db.Client
}

func NewPlaneRepositoryImpl(client *db.Client) ports.PlaneRepository {
	return &PlaneRepositoryImpl{client: client}
}

func (r *PassengerRepositoryImpl) GetAllPlanes() ([]entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying all planes", PASSENGER_LOG_PREFIX))

	ctx := context.Background()
	ref := r.client.NewRef("planes")

	var plane []entities.Plane
	err := ref.Get(ctx, &plane)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying all planes: %v", PASSENGER_LOG_PREFIX, err))
		return []entities.Plane{}, err
	}

	return plane, nil
}

func (r *PassengerRepositoryImpl) AddPlane(request entities.AddPlaneRequest) error {
	// Will be added
	return nil
}

func (r *PassengerRepositoryImpl) SetplaneStatus(request entities.SetPlaneStatusRequest) error {
	// Will be added
	return nil
}

func (s *PassengerRepositoryImpl) GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error) {
	// Will be added
	return entities.Plane{}, nil
}

func (s *PassengerRepositoryImpl) GetPlaneByFlightNumber(request entities.GetPlaneByFlightNumberRequest) (entities.Plane, error) {
	// Will be added
	return entities.Plane{}, nil
}

func (s *PassengerRepositoryImpl) GetPlaneByLocation(request entities.GetPlaneByLocationRequest) ([]entities.Plane, error) {
	// Will be added
	return []entities.Plane{}, nil
}
