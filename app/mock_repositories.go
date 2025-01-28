package app

import (
	"ams-service/core/entities"
)

type MockUserRepository struct{}

func (m *MockUserRepository) RegisterUser(user entities.User) error {
	return nil
}

type MockPassengerRepository struct{}

func (m *MockPassengerRepository) GetPassengerByID(passengerID string) (entities.Passenger, error) {
	return entities.Passenger{
		NationalId: "12345678901",
		PnrNo:      "ABC123",
		Name:       "John",
		Surname:    "Doe",
	}, nil
}

func (m *MockPassengerRepository) OnlineCheckInPassenger(pnr, surname string) error {
	return nil
}
