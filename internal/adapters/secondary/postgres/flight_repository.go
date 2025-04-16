package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"database/sql"

	"github.com/rs/zerolog/log"
)

type FlightRepositoryImpl struct {
	db *sql.DB
}

func NewFlightRepositoryImpl(db *sql.DB) secondary.FlightRepository {
	return &FlightRepositoryImpl{db: db}
}

// Helper function to scan flight rows
func scanFlightRow(row *sql.Row) (entities.Flight, error) {
	var flight entities.Flight
	err := row.Scan(
		&flight.FlightNumber,
		&flight.DepartureAirport,
		&flight.DestinationAirport,
		&flight.DepartureDateTime,
		&flight.ArrivalDateTime,
		&flight.DepartureGateNumber,
		&flight.DestinationGateNumber,
		&flight.PlaneRegistration,
		&flight.Status,
		&flight.Price,
	)
	return flight, err
}

// Helper function to scan multiple flight rows
func scanFlightRows(rows *sql.Rows) ([]entities.Flight, error) {
	var flights []entities.Flight
	for rows.Next() {
		var flight entities.Flight
		err := rows.Scan(
			&flight.FlightNumber,
			&flight.DepartureAirport,
			&flight.DestinationAirport,
			&flight.DepartureDateTime,
			&flight.ArrivalDateTime,
			&flight.DepartureGateNumber,
			&flight.DestinationGateNumber,
			&flight.PlaneRegistration,
			&flight.Status,
			&flight.Price,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning flight row")
			return nil, err
		}
		flights = append(flights, flight)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over flight rows")
		return nil, err
	}
	return flights, nil
}

func (r *FlightRepositoryImpl) GetSpecificFlight(request entities.GetSpecificFlightRequest) (entities.Flight, error) {
	log.Info().Str("flight_number", request.FlightNumber).Str("departure_datetime", request.DepartureDateTime).Msg("Getting flight by number and departure datetime")

	query := `
        SELECT flight_number, departure_airport, destination_airport, departure_datetime, arrival_datetime, 
               departure_gate_number, destination_gate_number, plane_registration, status, price 
        FROM flights 
        WHERE flight_number = $1 AND departure_datetime = $2`
	row := r.db.QueryRow(query, request.FlightNumber, request.DepartureDateTime)
	flight, err := scanFlightRow(row)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Str("departure_datetime", request.DepartureDateTime).Msg("Error getting flight by number and departure datetime")
		return entities.Flight{}, err
	}

	return flight, nil
}

func (r *FlightRepositoryImpl) GetAllFlights() ([]entities.Flight, error) {
	log.Info().Msg("Getting all flights")

	query := `
        SELECT flight_number, departure_airport, destination_airport, departure_datetime, arrival_datetime, 
               departure_gate_number, destination_gate_number, plane_registration, status, price 
        FROM flights`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Error().Err(err).Msg("Error querying all flights")
		return nil, err
	}
	defer rows.Close()

	return scanFlightRows(rows)
}

func (r *FlightRepositoryImpl) GetAllFlightsDestinationDateFlights(request entities.GetAllFlightsDestinationDateRequest) ([]entities.Flight, error) {
	log.Info().
		Str("departure_airport", request.DepartureAirport).
		Str("destination_airport", request.DestinationAirport).
		Str("departure_datetime", request.DepartureDateTime).
		Msg("Querying specific flights")

	query := `
        SELECT flight_number, departure_airport, destination_airport, departure_datetime, arrival_datetime, 
               departure_gate_number, destination_gate_number, plane_registration, status, price 
        FROM flights 
        WHERE departure_airport = $1 AND destination_airport = $2 AND departure_datetime::date = $3`
	rows, err := r.db.Query(query, request.DepartureAirport, request.DestinationAirport, request.DepartureDateTime)
	if err != nil {
		log.Error().Err(err).Msg("Error querying specific flights")
		return nil, err
	}
	defer rows.Close()

	return scanFlightRows(rows)
}

func (r *FlightRepositoryImpl) GetAllActiveFlights() ([]entities.Flight, error) {
	log.Info().Msg("Querying all active flights")

	query := `
        SELECT flight_number, departure_airport, destination_airport, departure_datetime, arrival_datetime, 
               departure_gate_number, destination_gate_number, plane_registration, status, price 
        FROM flights 
        WHERE status = 'scheduled'`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Error().Err(err).Msg("Error querying all active flights")
		return nil, err
	}
	defer rows.Close()

	return scanFlightRows(rows)
}

func (r *FlightRepositoryImpl) CancelFlight(request entities.CancelFlightRequest) error {
	log.Info().Str("flight_number", request.FlightNumber).Str("flight_date", request.FlightDate).Msg("Canceling flight")

	query := `UPDATE flights SET status = 'cancelled' WHERE flight_number = $1 AND departure_datetime::date = $2`
	result, err := r.db.Exec(query, request.FlightNumber, request.FlightDate)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Str("flight_date", request.FlightDate).Msg("Error canceling flight")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving rows affected")
		return err
	}

	if rowsAffected == 0 {
		log.Warn().Str("flight_number", request.FlightNumber).Str("flight_date", request.FlightDate).Msg("No flight found to cancel")
		return sql.ErrNoRows
	}

	return nil
}

func (r *FlightRepositoryImpl) AddFlight(request entities.AddFlightRequest) error {
	for _, flight := range request {
		query := `
        INSERT INTO flights (
            flight_number, departure_airport, destination_airport, departure_datetime, arrival_datetime,
            departure_gate_number, destination_gate_number, plane_registration, status, price
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    `
		_, err := r.db.Exec(
			query,
			flight.FlightNumber,
			flight.DepartureAirport,
			flight.DestinationAirport,
			flight.DepartureDateTime,
			flight.ArrivalDateTime,
			flight.DepartureGateNumber,
			flight.DestinationGateNumber,
			flight.PlaneRegistration,
			flight.Status,
			flight.Price,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error adding flight")
			return err
		}
	}

	return nil
}
