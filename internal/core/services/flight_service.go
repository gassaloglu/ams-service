package services

import (
	"ams-service/internal/core/entities"

	"github.com/rs/zerolog/log"
)

type FlightRepository interface {
	GetSpecificFlight(request entities.GetSpecificFlightRequest) (entities.Flight, error)
	GetAllFlights() ([]entities.Flight, error)
	GetAllSpecificFlights(request entities.GetSpecificFlightsRequest) ([]entities.Flight, error)
	GetAllActiveFlights() ([]entities.Flight, error)
	CancelFlight(request entities.CancelFlightRequest) error
}

type FlightService struct {
	repo FlightRepository
}

func NewFlightService(repo FlightRepository) *FlightService {
	return &FlightService{repo: repo}
}

func (s *FlightService) GetSpecificFlight(request entities.GetSpecificFlightRequest, userID string, resultChan chan<- entities.Flight, errorChan chan<- error) {
	go func() {
		flight, err := s.repo.GetSpecificFlight(request)
		if err != nil {
			log.Error().Err(err).Str("user_id", userID).Msg("Error getting flight by number and departure datetime")
			errorChan <- err
			return
		}
		log.Info().Str("user_id", userID).Msg("Successfully retrieved flight by number and departure datetime")
		resultChan <- flight
	}()
}

func (s *FlightService) GetAllFlights() ([]entities.Flight, error) {
	flights, err := s.repo.GetAllFlights()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all flights")
		return nil, err
	}
	log.Info().Msg("Successfully retrieved all flights")
	return flights, nil
}

func (s *FlightService) GetAllSpecificFlights(request entities.GetSpecificFlightsRequest) ([]entities.Flight, error) {
	flights, err := s.repo.GetAllSpecificFlights(request)
	if err != nil {
		log.Error().Err(err).Msg("Error getting specific flights")
		return nil, err
	}
	log.Info().Msg("Successfully retrieved specific flights")
	return flights, nil
}

func (s *FlightService) GetAllActiveFlights() ([]entities.Flight, error) {
	flights, err := s.repo.GetAllActiveFlights()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all active flights")
		return nil, err
	}
	log.Info().Msg("Successfully retrieved all active flights")
	return flights, nil
}

func (s *FlightService) CancelFlight(request entities.CancelFlightRequest) error {
	err := s.repo.CancelFlight(request)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error canceling flight")
		return err
	}
	log.Info().Str("flight_number", request.FlightNumber).Msg("Successfully canceled flight")
	return nil
}
