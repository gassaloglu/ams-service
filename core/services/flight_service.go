package services

import (
	"ams-service/core/entities"

	"github.com/rs/zerolog/log"
)

var FLIGHT_LOG_PREFIX string = "flight_service.go"

type FlightRepository interface {
	GetSpecificFlight(request entities.GetSpecificFlightRequest) (entities.Flight, error)
	GetAllFlights() ([]entities.Flight, error)
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
