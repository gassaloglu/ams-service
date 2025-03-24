package services

import (
	"ams-service/core/entities"
	"ams-service/middlewares"
	"fmt"
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
			middlewares.LogError(fmt.Sprintf("%s - Error getting flight by number and departure datetime for user %s: %v", FLIGHT_LOG_PREFIX, userID, err))
			errorChan <- err
			return
		}
		middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved flight by number and departure datetime for user %s", FLIGHT_LOG_PREFIX, userID))
		resultChan <- flight
	}()
}

func (s *FlightService) GetAllFlights() ([]entities.Flight, error) {
	flights, err := s.repo.GetAllFlights()
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting all flights: %v", FLIGHT_LOG_PREFIX, err))
		return nil, err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved all flights", FLIGHT_LOG_PREFIX))
	return flights, nil
}
