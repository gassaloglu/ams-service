package services

import (
	"ams-service/core/entities"
	"ams-service/middlewares"
	"fmt"
)

var PASSENGER_LOG_PREFIX string = "passenger_service.go"

type PassengerRepository interface {
	GetPassengerByID(passengerID string) (entities.Passenger, error)
	OnlineCheckInPassenger(pnr string, surname string) error
}

type PassengerService struct {
	repo PassengerRepository
}

func NewPassengerService(repo PassengerRepository) *PassengerService {
	return &PassengerService{repo: repo}
}

func (s *PassengerService) GetPassengerByID(passengerID string) (entities.Passenger, error) {
	passenger, err := s.repo.GetPassengerByID(passengerID)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting passenger by ID %s: %v", PASSENGER_LOG_PREFIX, passengerID, err))
		return entities.Passenger{}, err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved passenger by ID %s", PASSENGER_LOG_PREFIX, passengerID))
	return passenger, nil
}

func (s *PassengerService) OnlineCheckInPassenger(pnr string, surname string) error {
	err := s.repo.OnlineCheckInPassenger(pnr, surname)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error checking in passenger with PNR %s and surname %s: %v", PASSENGER_LOG_PREFIX, pnr, surname, err))
		return err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully checked in passenger with PNR %s and surname %s", PASSENGER_LOG_PREFIX, pnr, surname))
	return nil
}
