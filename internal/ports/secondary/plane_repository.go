package secondary

import (
	"ams-service/internal/core/entities"
)

type PlaneRepository interface {
	FindAll(request *entities.GetAllPlanesRequest) ([]entities.Plane, error)
	Create(plane *entities.Plane) (*entities.Plane, error)
}
