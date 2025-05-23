package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type PassengerRepositoryImpl struct {
	db *gorm.DB
}

func NewPassengerRepositoryImpl(db *gorm.DB) secondary.PassengerRepository {
	db.AutoMigrate(&entities.Passenger{})
	return &PassengerRepositoryImpl{db: db}
}

func (r *PassengerRepositoryImpl) FindById(id uint) (*entities.Passenger, error) {
	var passenger entities.Passenger
	result := r.db.Where("id", id).First(&passenger)
	return &passenger, result.Error
}

func (r *PassengerRepositoryImpl) GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error) {
	var passenger entities.Passenger
	result := r.db.Where("national_id", request.NationalId).First(&passenger)
	return passenger, result.Error
}

func (r *PassengerRepositoryImpl) GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error) {
	var passenger entities.Passenger
	result := r.db.
		Where("LOWER(pnr_no) = ?", strings.ToLower(request.PNR)).
		Where("LOWER(surname) = ?", strings.ToLower(request.Surname)).
		Where("status = ?", "active").
		First(&passenger)
	return passenger, result.Error
}

func (r *PassengerRepositoryImpl) OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error {
	result := r.db.Model(&entities.Passenger{}).
		Where("LOWER(pnr_no) = ?", strings.ToLower(request.PNR)).
		Where("LOWER(surname) = ?", strings.ToLower(request.Surname)).
		Where("status = ?", "active").
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

func (r *PassengerRepositoryImpl) CreatePassenger(request *entities.Passenger) (*entities.Passenger, error) {
	clone := *request
	result := r.db.Create(&clone)
	return &clone, result.Error
}

func (r *PassengerRepositoryImpl) GetAllPassengers() ([]entities.Passenger, error) {
	var passengers []entities.Passenger
	result := r.db.Find(&passengers)
	return passengers, result.Error
}

func (r *PassengerRepositoryImpl) EmployeeCheckInPassenger(request entities.EmployeeCheckInRequest) (entities.Passenger, error) {
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

func (r *PassengerRepositoryImpl) FindPassengersMatchingAnyUniquePassengerInfo(p *entities.PassengerInfo) (*entities.Passenger, error) {
	var passenger entities.Passenger

	result := r.db.
		Joins("JOIN flights ON flights.id = passengers.flight_id").
		Where("flights.flight_number = ?", p.FlightNumber).
		Where(
			r.db.
				Where("national_id = ?", p.NationalID).
				Or("email = ?", p.Email).
				Or("phone = ?", p.Phone).
				Or("seat = ?", p.Seat),
		).
		First(&passenger)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no passenger found with the given ID")
	}

	return &passenger, nil
}
