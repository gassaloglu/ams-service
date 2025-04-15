package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"
)

type PlaneService struct {
	repo secondary.PlaneRepository
}

func NewPlaneService(repo secondary.PlaneRepository) primary.PlaneService {
	return &PlaneService{repo: repo}
}

func (s *PlaneService) GetAllPlanes(request entities.GetAllPlanesRequest) ([]entities.Plane, error) {
	return s.repo.GetAllPlanes(request)
}

func (s *PlaneService) GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error) {
	return s.repo.GetPlaneByRegistration(request)
}

func (s *PlaneService) AddPlane(request entities.AddPlaneRequest) error {
	return s.repo.AddPlane(request)
}

func (s *PlaneService) SetPlaneStatus(request entities.SetPlaneStatusRequest) error {
	return s.repo.SetPlaneStatus(request)
}
