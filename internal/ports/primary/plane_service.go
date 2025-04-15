package primary

import (
	"ams-service/internal/core/entities"
)

type PlaneService interface {
	GetAllPlanes(entities.GetAllPlanesRequest) ([]entities.Plane, error)
	GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error)
	AddPlane(request entities.AddPlaneRequest) error
	SetPlaneStatus(request entities.SetPlaneStatusRequest) error
}
