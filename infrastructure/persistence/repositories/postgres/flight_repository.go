package postgres

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"database/sql"
	"fmt"
)

var FLIGHT_LOG_PREFIX string = "flight_repository.go"

type FlightRepositoryImpl struct {
	db *sql.DB
}

func NewFlightRepositoryImpl(db *sql.DB) ports.FlightRepository {
	return &FlightRepositoryImpl{db: db}
}

func (r *FlightRepositoryImpl) GetSpecificFlight(request entities.GetSpecificFlightRequest) (entities.Flight, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Getting flight by number: %s and departure datetime: %s", FLIGHT_LOG_PREFIX, request.FlightNumber, request.DepartureDateTime))

	query := `SELECT flight_number, departure_airport, destination_airport, departure_datetime, arrival_datetime, departure_gate_number, destination_gate_number, plane_registration, status, price FROM flights WHERE flight_number = $1 AND departure_datetime = $2`
	row := r.db.QueryRow(query, request.FlightNumber, request.DepartureDateTime)

	var flight entities.Flight
	err := row.Scan(&flight.FlightNumber, &flight.DepartureAirport, &flight.DestinationAirport, &flight.DepartureDateTime, &flight.ArrivalDateTime, &flight.DepartureGateNumber, &flight.DestinationGateNumber, &flight.PlaneRegistration, &flight.Status, &flight.Price)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting flight by number and departure datetime: %v", FLIGHT_LOG_PREFIX, err))
		return entities.Flight{}, err
	}

	return flight, nil
}

func (r *FlightRepositoryImpl) GetAllFlights() ([]entities.Flight, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Getting all flights", FLIGHT_LOG_PREFIX))

	query := `SELECT flight_number, departure_airport, destination_airport, departure_datetime, arrival_datetime, departure_gate_number, destination_gate_number, plane_registration, status, price FROM flights`
	rows, err := r.db.Query(query)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying all flights: %v", FLIGHT_LOG_PREFIX, err))
		return nil, err
	}
	defer rows.Close()

	var flights []entities.Flight
	for rows.Next() {
		var flight entities.Flight
		err := rows.Scan(&flight.FlightNumber, &flight.DepartureAirport, &flight.DestinationAirport, &flight.DepartureDateTime, &flight.ArrivalDateTime, &flight.DepartureGateNumber, &flight.DestinationGateNumber, &flight.PlaneRegistration, &flight.Status, &flight.Price)
		if err != nil {
			middlewares.LogError(fmt.Sprintf("%s - Error scanning flight: %v", FLIGHT_LOG_PREFIX, err))
			return nil, err
		}
		flights = append(flights, flight)
	}

	return flights, nil
}
