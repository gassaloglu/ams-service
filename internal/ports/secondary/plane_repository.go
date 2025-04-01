package secondary

import (
	"ams-service/internal/core/entities"
)

type PlaneRepository interface {
	GetAllPlanes() ([]entities.Plane, error)
	GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error)
	GetPlaneByFlightNumber(request entities.GetPlaneByFlightNumberRequest) (entities.Plane, error)
	GetPlaneByLocation(request entities.GetPlaneByLocationRequest) ([]entities.Plane, error)
	AddPlane(request entities.AddPlaneRequest) error
	SetPlaneStatus(request entities.SetPlaneStatusRequest) error
}
