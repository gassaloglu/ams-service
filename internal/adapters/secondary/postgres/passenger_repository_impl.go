package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"errors"

	"gorm.io/gorm"
)

type PassengerRepositoryImpl struct {
	db *gorm.DB
}

func NewPassengerRepositoryImpl(db *gorm.DB) secondary.PassengerRepository {
	db.AutoMigrate(&entities.Passenger{})
	return &PassengerRepositoryImpl{db: db}
}

func (r *PassengerRepositoryImpl) GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error) {
	var passenger entities.Passenger
	result := r.db.Where("national_id", request.NationalId).Find(&passenger)
	return passenger, result.Error
}

func (r *PassengerRepositoryImpl) GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error) {
	var passenger entities.Passenger
	result := r.db.
		Where("pnr", request.PNR).
		Where("surname", request.Surname).
		Find(&passenger)
	return passenger, result.Error
}

func (r *PassengerRepositoryImpl) OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error {
	result := r.db.Model(&entities.Passenger{}).
		Where("pnr", request.PNR).
		Where("surname", request.Surname).
		Update("check_in", true)

	return result.Error
}

func (r *PassengerRepositoryImpl) GetPassengersBySpecificFlight(request entities.GetPassengersBySpecificFlightRequest) ([]entities.Passenger, error) {
	var passengers []entities.Passenger
	var flight entities.Flight
	result := r.db.Model(&flight).
		Where("flight_number", request.FlightNumber).
		Where("departure_datetime", request.DepartureDateTime)

	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.Model(&passengers).Where("flight_id", flight.ID)
	return passengers, result.Error
}

func (r *PassengerRepositoryImpl) CreatePassenger(request entities.CreatePassengerRequest) error {
	result := r.db.Create(request.Passenger)
	return result.Error
}

func (r *PassengerRepositoryImpl) GetAllPassengers() ([]entities.Passenger, error) {
	var passengers []entities.Passenger
	result := r.db.Find(&passengers)
	return passengers, result.Error
}

func (r *PassengerRepositoryImpl) EmployeeCheckInPassenger(request entities.EmployeeCheckInRequest) (entities.Passenger, error) {
	// TODO
	return entities.Passenger{}, errors.ErrUnsupported
}

func (r *PassengerRepositoryImpl) CancelPassenger(request entities.CancelPassengerRequest) error {
	result := r.db.Model(&entities.Passenger{}).
		Where("id = ?", request.PassengerID).
		Update("status", "inactive")

	if result.RowsAffected == 0 {
		return errors.New("no passenger found with the given ID")
	}

	return result.Error
}
