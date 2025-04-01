package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"

	"github.com/rs/zerolog/log"
)

type PlaneService struct {
	repo secondary.PlaneRepository
}

func NewPlaneService(repo secondary.PlaneRepository) primary.PlaneService {
	return &PlaneService{repo: repo}
}

func (s *PlaneService) GetAllPlanes() ([]entities.Plane, error) {
	planes, err := s.repo.GetAllPlanes()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all planes")
		return nil, err
	}
	log.Info().Msg("Successfully retrieved all planes")
	return planes, nil
}

func (s *PlaneService) GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error) {
	plane, err := s.repo.GetPlaneByRegistration(request)
	if err != nil {
		log.Error().Err(err).Str("registration", request.PlaneRegistration).Msg("Error getting plane by registration")
		return entities.Plane{}, err
	}
	log.Info().Str("registration", request.PlaneRegistration).Msg("Successfully retrieved plane by registration")
	return plane, nil
}

func (s *PlaneService) GetPlaneByFlightNumber(request entities.GetPlaneByFlightNumberRequest) (entities.Plane, error) {
	plane, err := s.repo.GetPlaneByFlightNumber(request)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error getting plane by flight number")
		return entities.Plane{}, err
	}
	log.Info().Str("flight_number", request.FlightNumber).Msg("Successfully retrieved plane by flight number")
	return plane, nil
}

func (s *PlaneService) GetPlaneByLocation(request entities.GetPlaneByLocationRequest) ([]entities.Plane, error) {
	planes, err := s.repo.GetPlaneByLocation(request)
	if err != nil {
		log.Error().Err(err).Str("location", request.Location).Msg("Error getting planes by location")
		return nil, err
	}
	log.Info().Str("location", request.Location).Msg("Successfully retrieved planes by location")
	return planes, nil
}

func (s *PlaneService) AddPlane(request entities.AddPlaneRequest) error {
	err := s.repo.AddPlane(request)
	if err != nil {
		log.Error().Err(err).Msg("Error adding plane")
		return err
	}
	log.Info().Msg("Successfully added plane")
	return nil
}

func (s *PlaneService) SetPlaneStatus(request entities.SetPlaneStatusRequest) error {
	err := s.repo.SetPlaneStatus(request)
	if err != nil {
		log.Error().Err(err).Msg("Error setting plane status")
		return err
	}
	log.Info().Msg("Successfully set plane status")
	return nil
}
