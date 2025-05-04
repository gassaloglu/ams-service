package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"

	"gorm.io/gorm"
)

type FlightRepositoryImpl struct {
	db *gorm.DB
}

func NewFlightRepositoryImpl(db *gorm.DB) secondary.FlightRepository {
	db.AutoMigrate(&entities.Flight{})
	return &FlightRepositoryImpl{db: db}
}

func (r *FlightRepositoryImpl) FindById(id string) (*entities.Flight, error) {
	var flight entities.Flight

	result := r.db.
		Where("id", id).
		Find(&flight)

	return &flight, result.Error
}

func (r *FlightRepositoryImpl) FindAll() ([]entities.Flight, error) {
	var flights []entities.Flight
	result := r.db.Find(&flights)
	return flights, result.Error
}

func (r *FlightRepositoryImpl) FindAllActive() ([]entities.Flight, error) {
	var flights []entities.Flight
	result := r.db.Where("status", "scheduled").Or("status", "delayed").Find(&flights)
	return flights, result.Error
}

func (r *FlightRepositoryImpl) Create(flight *entities.Flight) error {
	result := r.db.Create(flight)
	return result.Error
}

func (r *FlightRepositoryImpl) CreateAll(flights []entities.Flight) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		result := r.db.Create(flights)
		return result.Error
	})
}

func (r *FlightRepositoryImpl) FindSeatsByFlightId(id string) ([]int, error) {
	var seats []int

	result := r.db.Model(&entities.Passenger{}).
		Select("seat").
		Where("id", id).
		Where("status", "scheduled").
		Or("status", "delayed").
		Find(&seats)

	return seats, result.Error
}
