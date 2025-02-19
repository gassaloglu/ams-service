package postgres

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"database/sql"
)

var PLANE_LOG_PREFIX string = "passenger_repository.go"

type PlaneRepositoryImpl struct {
	db *sql.DB
}

func NewPlaneRepositoryImpl(db *sql.DB) ports.PlaneRepository {
	return &PlaneRepositoryImpl{db: db}
}

func (r *PlaneRepositoryImpl) GetAllPlanes() ([]entities.Plane, error) {
	// Will be added
	return nil, nil
}

func (r *PlaneRepositoryImpl) AddPlane(request entities.AddPlaneRequest) error {
	// Will be added
	return nil
}

func (r *PlaneRepositoryImpl) SetPlaneStatus(request entities.SetPlaneStatusRequest) error {
	// Will be added
	return nil
}

func (s *PlaneRepositoryImpl) GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error) {
	// Will be added
	return entities.Plane{}, nil
}

func (s *PlaneRepositoryImpl) GetPlaneByFlightNumber(request entities.GetPlaneByFlightNumberRequest) (entities.Plane, error) {
	// Will be added
	return entities.Plane{}, nil
}

func (s *PlaneRepositoryImpl) GetPlaneByLocation(request entities.GetPlaneByLocationRequest) ([]entities.Plane, error) {
	// Will be added
	return []entities.Plane{}, nil
}
