package services

import (
	"ams-service/core/entities"

	"github.com/rs/zerolog/log"
)

var PASSENGER_LOG_PREFIX string = "passenger_service.go"

type PassengerRepository interface {
	GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error)
	GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error)
	OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error
	GetPassengersBySpecificFlight(request entities.GetPassengersBySpecificFlightRequest) ([]entities.Passenger, error)
	CreatePassenger(request entities.CreatePassengerRequest) error
	GetAllPassengers() ([]entities.Passenger, error)
}

type PassengerService struct {
	repo PassengerRepository
}

func NewPassengerService(repo PassengerRepository) *PassengerService {
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
