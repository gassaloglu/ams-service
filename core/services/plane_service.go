package services

import (
	"ams-service/core/entities"
	"ams-service/middlewares"
	"fmt"
)

var PLANE_LOG_PREFIX string = "plane_service.go"

type PlaneRepository interface {
	GetAllPlanes() ([]entities.Plane, error)
	GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error)
	GetPlaneByFlightNumber(request entities.GetPlaneByFlightNumberRequest) (entities.Plane, error)
	GetPlaneByLocation(request entities.GetPlaneByLocationRequest) ([]entities.Plane, error)
	AddPlane(request entities.AddPlaneRequest) error
	SetPlaneStatus(request entities.SetPlaneStatusRequest) error
}

type PlaneService struct {
	repo PlaneRepository
}

func NewPlaneService(repo PlaneRepository) *PlaneService {
	return &PlaneService{repo: repo}
}

func (s *PlaneService) GetAllPlanes() ([]entities.Plane, error) {
	planes, err := s.repo.GetAllPlanes()
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting all planes: %v", PLANE_LOG_PREFIX, err))
		return nil, err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved all planes", PLANE_LOG_PREFIX))
	return planes, nil
}

func (s *PlaneService) GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error) {
	plane, err := s.repo.GetPlaneByRegistration(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting plane by registration %s: %v", PLANE_LOG_PREFIX, request.PlaneRegistration, err))
		return entities.Plane{}, err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved plane by registration %s", PLANE_LOG_PREFIX, request.PlaneRegistration))
	return plane, nil
}

func (s *PlaneService) GetPlaneByFlightNumber(request entities.GetPlaneByFlightNumberRequest) (entities.Plane, error) {
	plane, err := s.repo.GetPlaneByFlightNumber(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting plane by flight number %s: %v", PLANE_LOG_PREFIX, request.FlightNumber, err))
		return entities.Plane{}, err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved plane by flight number %s", PLANE_LOG_PREFIX, request.FlightNumber))
	return plane, nil
}

func (s *PlaneService) GetPlaneByLocation(request entities.GetPlaneByLocationRequest) ([]entities.Plane, error) {
	planes, err := s.repo.GetPlaneByLocation(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting planes by location %s: %v", PLANE_LOG_PREFIX, request.Location, err))
		return nil, err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved planes by location %s", PLANE_LOG_PREFIX, request.Location))
	return planes, nil
}

func (s *PlaneService) AddPlane(request entities.AddPlaneRequest) error {
	err := s.repo.AddPlane(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error adding plane: %v", PLANE_LOG_PREFIX, err))
		return err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully added plane", PLANE_LOG_PREFIX))
	return nil
}

func (s *PlaneService) SetPlaneStatus(request entities.SetPlaneStatusRequest) error {
	err := s.repo.SetPlaneStatus(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error setting plane status: %v", PLANE_LOG_PREFIX, err))
		return err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully set plane status", PLANE_LOG_PREFIX))
	return nil
}
