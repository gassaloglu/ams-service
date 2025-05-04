package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"

	"github.com/sourcegraph/conc/iter"
)

type PlaneService struct {
	repo secondary.PlaneRepository
}

func NewPlaneService(repo secondary.PlaneRepository) primary.PlaneService {
	return &PlaneService{repo: repo}
}

func (s *PlaneService) FindAll(request *entities.GetAllPlanesRequest) ([]entities.Plane, error) {
	return s.repo.FindAll(request)
}

func (s *PlaneService) Create(request *entities.CreatePlaneRequest) (*entities.Plane, error) {
	plane := mapCreatePlaneRequestToPlaneEntity(request)
	return s.repo.Create(&plane)
}

func (s *PlaneService) CreateAll(requests []entities.CreatePlaneRequest) error {
	planes := iter.Map(requests, mapCreatePlaneRequestToPlaneEntity)
	return s.repo.CreateAll(planes)
}

func mapCreatePlaneRequestToPlaneEntity(request *entities.CreatePlaneRequest) entities.Plane {
	return entities.Plane{
		Registration: request.Registration,
		Model:        request.Model,
		Manufacturer: request.Manufacturer,
		Capacity:     request.Capacity,
		Status:       request.Status,
	}
}
