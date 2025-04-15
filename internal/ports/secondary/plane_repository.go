package secondary

import (
	"ams-service/internal/core/entities"
)

type PlaneRepository interface {
	GetAllPlanes(request entities.GetAllPlanesRequest) ([]entities.Plane, error)
	GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error)
	AddPlane(request entities.AddPlaneRequest) error
	SetPlaneStatus(request entities.SetPlaneStatusRequest) error
}
