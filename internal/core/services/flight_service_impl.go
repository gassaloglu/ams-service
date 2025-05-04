package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"

	"github.com/rs/zerolog/log"
)

type FlightService struct {
	repo secondary.FlightRepository
}

func NewFlightService(repo secondary.FlightRepository) primary.FlightService {
	return &FlightService{repo: repo}
}

func (s *FlightService) GetSpecificFlight(request entities.GetSpecificFlightRequest) (entities.Flight, error) {
	flight, err := s.repo.GetSpecificFlight(request)
	if err != nil {
		log.Error().Err(err).Msg("Error getting flight by number and departure datetime")
		return entities.Flight{}, err
	}
	log.Info().Msg("Successfully retrieved flight by number and departure datetime")
	return flight, nil
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

func (s *FlightService) GetAllFlightsDestinationDateFlights(request entities.GetAllFlightsDestinationDateRequest) ([]entities.Flight, error) {
	flights, err := s.repo.GetAllFlightsDestinationDateFlights(request)
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
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("TODO: Error canceling flight")
		return err
	}
	log.Info().Str("flight_number", request.FlightNumber).Msg("Successfully canceled flight")
	return nil
}

func (s *FlightService) AddFlight(request entities.AddFlightRequest) error {
	err := s.repo.AddFlight(request)
	if err != nil {
		log.Error().Err(err).Msg("Error adding flight")
		return err
	}
	log.Info().Msg("Successfully added flight")
	return nil
}

func (s *FlightService) FetchSeatMap(request entities.FetchSeatMapRequest) ([]int, error) {
	seats, err := s.repo.FetchSeatMap(request)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightID).Msg("Error fetching seat map")
		return nil, err
	}
	return seats, nil
}
