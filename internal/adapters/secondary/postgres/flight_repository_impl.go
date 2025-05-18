package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"ams-service/internal/utils"
	"time"

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

func (r *FlightRepositoryImpl) FindByFlightNumber(flightNumber string) (*entities.Flight, error) {
	var flight entities.Flight

	result := r.db.
		Where("flight_number", flightNumber).
		Find(&flight)

	return &flight, result.Error
}

func (r *FlightRepositoryImpl) FindAll(request *entities.GetAllFlightsRequest) ([]entities.Flight, error) {
	var flights []entities.Flight
	result := buildFindAllFlightsQuery(r.db, request).Find(&flights)
	return flights, result.Error
}

func (r *FlightRepositoryImpl) FindAllActive(request *entities.GetAllFlightsRequest) ([]entities.Flight, error) {
	var flights []entities.Flight

	result := buildFindAllFlightsQuery(r.db, request).
		Where("status", "scheduled").
		Find(&flights)

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
		Joins("JOIN flights ON flights.id = passengers.flight_id").
		Where("passengers.status = ?", "active").
		Where("flights.status = ?", "scheduled").
		Where("flights.id = ?", id).
		Order("seat").
		Find(&seats)

	return seats, result.Error
}

func buildFindAllFlightsQuery(db *gorm.DB, request *entities.GetAllFlightsRequest) *gorm.DB {
	if len(request.ID) > 0 {
		db = db.Where("id IN ?", request.ID)
	}

	if len(request.FlightNumber) > 0 {
		db = db.Where("flight_number IN ?", request.FlightNumber)
	}

	if len(request.DepartureAirport) > 0 {
		db = db.Where("departure_airport IN ?", request.DepartureAirport)
	}

	if len(request.DestinationAirport) > 0 {
		db = db.Where("destination_airport IN ?", request.DestinationAirport)
	}

	if request.DepartureDatetime != nil {
		db = utils.BuildComparableQueryForField[time.Time](db, request.DepartureDatetime, "departure_datetime")
	}

	if len(request.Status) > 0 {
		db = db.Where("status IN ?", request.Status)
	}

	if request.Price != nil {
		db = utils.BuildComparableQueryForField[float64](db, request.Price, "price")
	}

	return db
}
