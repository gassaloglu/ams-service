package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"

	"github.com/rs/zerolog/log"
)

type PassengerService struct {
	repo secondary.PassengerRepository
}

func NewPassengerService(repo secondary.PassengerRepository) primary.PassengerService {
	return &PassengerService{repo: repo}
}

func (s *PassengerService) GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error) {
	passenger, err := s.repo.GetPassengerByID(request)
	if err != nil {
		log.Error().Err(err).Str("national_id", request.NationalId).Msg("Error getting passenger by ID")
		return entities.Passenger{}, err
	}
	log.Info().Str("national_id", request.NationalId).Msg("Successfully retrieved passenger by ID")
	return passenger, nil
}

func (s *PassengerService) GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error) {
	passenger, err := s.repo.GetPassengerByPNR(request)
	if err != nil {
		log.Error().Err(err).Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Error getting passenger by PNR")
		return entities.Passenger{}, err
	}
	log.Info().Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Successfully retrieved passenger by PNR")
	return passenger, nil
}

func (s *PassengerService) OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error {
	err := s.repo.OnlineCheckInPassenger(request)
	if err != nil {
		log.Error().Err(err).Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Error checking in passenger")
		return err
	}
	log.Info().Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Successfully checked in passenger")
	return nil
}

func (s *PassengerService) GetPassengersBySpecificFlight(request entities.GetPassengersBySpecificFlightRequest) ([]entities.Passenger, error) {
	passengers, err := s.repo.GetPassengersBySpecificFlight(request)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error getting passengers by specific flight")
		return nil, err
	}
	log.Info().Str("flight_number", request.FlightNumber).Msg("Successfully retrieved passengers by specific flight")
	return passengers, nil
}

func (s *PassengerService) CreatePassenger(request entities.CreatePassengerRequest) error {
	err := s.repo.CreatePassenger(request)
	if err != nil {
		log.Error().Err(err).Str("national_id", request.NationalId).Msg("Error creating passenger")
		return err
	}
	log.Info().Str("national_id", request.NationalId).Msg("Successfully created passenger")
	return nil
}

func (s *PassengerService) GetAllPassengers() ([]entities.Passenger, error) {
	passengers, err := s.repo.GetAllPassengers()
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving all passengers")
		return nil, err
	}
	log.Info().Msg("Successfully retrieved all passengers")
	return passengers, nil
}
