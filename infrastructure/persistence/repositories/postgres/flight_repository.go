package postgres

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"database/sql"

	"github.com/rs/zerolog/log"
)

var FLIGHT_LOG_PREFIX string = "flight_repository.go"

type FlightRepositoryImpl struct {
	db *sql.DB
}

func NewFlightRepositoryImpl(db *sql.DB) ports.FlightRepository {
	return &FlightRepositoryImpl{db: db}
}

func (r *FlightRepositoryImpl) GetSpecificFlight(request entities.GetSpecificFlightRequest) (entities.Flight, error) {
	log.Info().Str("flight_number", request.FlightNumber).Str("departure_datetime", request.DepartureDateTime).Msg("Getting flight by number and departure datetime")

	query := `SELECT flight_number, departure_airport, destination_airport, departure_datetime, arrival_datetime, departure_gate_number, destination_gate_number, plane_registration, status, price FROM flights WHERE flight_number = $1 AND departure_datetime = $2`
	row := r.db.QueryRow(query, request.FlightNumber, request.DepartureDateTime)

	var flight entities.Flight
	err := row.Scan(&flight.FlightNumber, &flight.DepartureAirport, &flight.DestinationAirport, &flight.DepartureDateTime, &flight.ArrivalDateTime, &flight.DepartureGateNumber, &flight.DestinationGateNumber, &flight.PlaneRegistration, &flight.Status, &flight.Price)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Str("departure_datetime", request.DepartureDateTime).Msg("Error getting flight by number and departure datetime")
		return entities.Flight{}, err
	}

	return flight, nil
}

func (r *FlightRepositoryImpl) GetAllFlights() ([]entities.Flight, error) {
	log.Info().Msg("Getting all flights")

	query := `SELECT flight_number, departure_airport, destination_airport, departure_datetime, arrival_datetime, departure_gate_number, destination_gate_number, plane_registration, status, price FROM flights`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Error().Err(err).Msg("Error querying all flights")
		return nil, err
	}
	defer rows.Close()

	var flights []entities.Flight
	for rows.Next() {
		var flight entities.Flight
		err := rows.Scan(&flight.FlightNumber, &flight.DepartureAirport, &flight.DestinationAirport, &flight.DepartureDateTime, &flight.ArrivalDateTime, &flight.DepartureGateNumber, &flight.DestinationGateNumber, &flight.PlaneRegistration, &flight.Status, &flight.Price)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning flight")
			return nil, err
		}
		flights = append(flights, flight)
	}

	return flights, nil
}

func (r *FlightRepositoryImpl) GetAllSpecificFlights(request entities.GetSpecificFlightsRequest) ([]entities.Flight, error) {
	log.Info().
		Str("departure_airport", request.DepartureAirport).
		Str("destination_airport", request.DestinationAirport).
		Str("departure_datetime", request.DepartureDateTime).
		Msg("Querying specific flights")

	query := `SELECT flight_number, departure_airport, destination_airport, departure_datetime, arrival_datetime, departure_gate_number, destination_gate_number, plane_registration, status, price 
              FROM flights 
              WHERE departure_airport = $1 AND destination_airport = $2 AND departure_datetime::date = $3`

	rows, err := r.db.Query(query, request.DepartureAirport, request.DestinationAirport, request.DepartureDateTime)
	if err != nil {
		log.Error().Err(err).Msg("Error querying specific flights")
		return nil, err
	}
	defer rows.Close()

	var flights []entities.Flight
	for rows.Next() {
		var flight entities.Flight
		err := rows.Scan(&flight.FlightNumber, &flight.DepartureAirport, &flight.DestinationAirport, &flight.DepartureDateTime, &flight.ArrivalDateTime, &flight.DepartureGateNumber, &flight.DestinationGateNumber, &flight.PlaneRegistration, &flight.Status, &flight.Price)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning flight row")
			return nil, err
		}
		flights = append(flights, flight)
	}

	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over flight rows")
		return nil, err
	}

	return flights, nil
}
