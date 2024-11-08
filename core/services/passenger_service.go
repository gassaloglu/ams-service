package services

import "ams-service/core/entities"

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
	return s.repo.GetPassengerByID(passengerID)
}

func (s *PassengerService) OnlineCheckInPassenger(pnr string, surname string) error {
	return s.repo.OnlineCheckInPassenger(pnr, surname)
}
