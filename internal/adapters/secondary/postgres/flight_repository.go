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

func (r *FlightRepositoryImpl) GetSpecificFlight(request entities.GetSpecificFlightRequest) (entities.Flight, error) {
	var flight entities.Flight

	result := r.db.
		Where("flight_number", request.FlightNumber).
		Where("departure_datetime", request.DepartureDateTime).
		Find(&flight)

	return flight, result.Error
}

func (r *FlightRepositoryImpl) GetAllFlights() ([]entities.Flight, error) {
	var flights []entities.Flight
	result := r.db.Find(&flights)
	return flights, result.Error
}

func (r *FlightRepositoryImpl) GetAllFlightsDestinationDateFlights(request entities.GetAllFlightsDestinationDateRequest) ([]entities.Flight, error) {
	var flights []entities.Flight

	result := r.db.Model(&entities.Flight{}).
		Where("departure_airport", request.DepartureAirport).
		Where("destination_airport", request.DestinationAirport).
		Where("departure_datetime", request.DepartureDateTime).
		Find(&flights)

	return flights, result.Error
}

func (r *FlightRepositoryImpl) GetAllActiveFlights() ([]entities.Flight, error) {
	var flights []entities.Flight
	result := r.db.Where("status", "scheduled").Find(&flights)
	return flights, result.Error
}

func (r *FlightRepositoryImpl) CancelFlight(request entities.CancelFlightRequest) error {
	result := r.db.Model(&entities.Flight{}).
		Where("flight_number", request.FlightNumber).
		Where("departure_datetime", request.FlightDate).
		Update("status", "cancelled")

	return result.Error
}

func (r *FlightRepositoryImpl) AddFlight(request entities.AddFlightRequest) error {
	result := r.db.Create(request)
	return result.Error
}

func (r *FlightRepositoryImpl) FetchSeatMap(request entities.FetchSeatMapRequest) ([]int, error) {
	var seats []int
	err := r.db.Model(&entities.Passenger{}).
		Select("seat").
		Where("flight_id", request.FlightID).
		Where("status", "active").
		Find(&seats).Error

	if err != nil {
		return nil, err
	}
	return seats, nil
}
