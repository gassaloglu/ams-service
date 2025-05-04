package primary

import (
	"ams-service/internal/core/entities"
)

type PlaneService interface {
	FindAll(*entities.GetAllPlanesRequest) ([]entities.Plane, error)
	Create(request *entities.CreatePlaneRequest) (*entities.Plane, error)
}
