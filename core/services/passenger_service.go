package services

import (
	"ams-service/core/entities"
	"ams-service/middlewares"
	"fmt"
)

var PASSENGER_LOG_PREFIX string = "passenger_service.go"

type PassengerRepository interface {
	GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error)
	GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error)
	OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error
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
		middlewares.LogError(fmt.Sprintf("%s - Error getting passenger by ID %s: %v", PASSENGER_LOG_PREFIX, request.NationalId, err))
		return entities.Passenger{}, err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved passenger by ID %s", PASSENGER_LOG_PREFIX, request.NationalId))
	return passenger, nil
}

func (s *PassengerService) GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error) {
	passenger, err := s.repo.GetPassengerByPNR(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting passenger by PNR %s and surname %s: %v", PASSENGER_LOG_PREFIX, request.PNR, request.Surname, err))
		return entities.Passenger{}, err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved passenger by PNR %s and surname %s", PASSENGER_LOG_PREFIX, request.PNR, request.Surname))
	return passenger, nil
}

func (s *PassengerService) OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error {
	err := s.repo.OnlineCheckInPassenger(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error checking in passenger with PNR %s and surname %s: %v", PASSENGER_LOG_PREFIX, request.PNR, request.Surname, err))
		return err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully checked in passenger with PNR %s and surname %s", PASSENGER_LOG_PREFIX, request.PNR, request.Surname))
	return nil
}
