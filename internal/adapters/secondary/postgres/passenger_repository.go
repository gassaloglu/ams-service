package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

type PassengerRepositoryImpl struct {
	db *sql.DB
}

func NewPassengerRepositoryImpl(db *sql.DB) secondary.PassengerRepository {
	return &PassengerRepositoryImpl{db: db}
}

func (r *PassengerRepositoryImpl) GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error) {
	log.Info().Str("national_id", request.NationalId).Msg("Querying passenger by ID")
	query := `SELECT id, national_id, pnr_no, flight_id, payment_id, baggage_allowance, baggage_id, fare_type, seat, meal, extra_baggage, check_in, name, surname, email, phone, gender, birth_date, cip_member, vip_member, disabled, child, created_at, updated_at FROM passengers WHERE national_id = $1`
	row := r.db.QueryRow(query, request.NationalId)

	var passenger entities.Passenger
	err := row.Scan(
		&passenger.ID,
		&passenger.NationalId,
		&passenger.PnrNo,
		&passenger.FlightId,
		&passenger.PaymentId,
		&passenger.BaggageAllowance,
		&passenger.BaggageId,
		&passenger.FareType,
		&passenger.Seat,
		&passenger.Meal,
		&passenger.ExtraBaggage,
		&passenger.CheckIn,
		&passenger.Name,
		&passenger.Surname,
		&passenger.Email,
		&passenger.Phone,
		&passenger.Gender,
		&passenger.BirthDate,
		&passenger.CipMember,
		&passenger.VipMember,
		&passenger.Disabled,
		&passenger.Child,
		&passenger.CreatedAt,
		&passenger.UpdatedAt,
	)
	if err != nil {
		log.Error().Err(err).Str("national_id", request.NationalId).Msg("Error querying passenger by ID")
		return entities.Passenger{}, err
	}

	return passenger, nil
}

func (r *PassengerRepositoryImpl) GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error) {
	log.Info().Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Querying passenger by PNR")
	query := `SELECT id, national_id, pnr_no, flight_id, payment_id, baggage_allowance, baggage_id, fare_type, seat, meal, extra_baggage, check_in, name, surname, email, phone, gender, birth_date, cip_member, vip_member, disabled, child, created_at, updated_at FROM passengers WHERE pnr_no = $1 AND surname = $2`
	row := r.db.QueryRow(query, request.PNR, request.Surname)

	var passenger entities.Passenger
	err := row.Scan(
		&passenger.ID,
		&passenger.NationalId,
		&passenger.PnrNo,
		&passenger.FlightId,
		&passenger.PaymentId,
		&passenger.BaggageAllowance,
		&passenger.BaggageId,
		&passenger.FareType,
		&passenger.Seat,
		&passenger.Meal,
		&passenger.ExtraBaggage,
		&passenger.CheckIn,
		&passenger.Name,
		&passenger.Surname,
		&passenger.Email,
		&passenger.Phone,
		&passenger.Gender,
		&passenger.BirthDate,
		&passenger.CipMember,
		&passenger.VipMember,
		&passenger.Disabled,
		&passenger.Child,
		&passenger.CreatedAt,
		&passenger.UpdatedAt,
	)
	if err != nil {
		log.Error().Err(err).Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Error querying passenger by PNR")
		return entities.Passenger{}, err
	}

	return passenger, nil
}

func (r *PassengerRepositoryImpl) OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error {
	log.Info().Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Checking in passenger")
	query := `UPDATE passengers SET check_in = true WHERE pnr_no = $1 AND surname = $2`
	_, err := r.db.Exec(query, request.PNR, request.Surname)
	if err != nil {
		log.Error().Err(err).Str("pnr", request.PNR).Str("surname", request.Surname).Msg("Error checking in passenger")
		return err
	}

	return nil
}

func (r *PassengerRepositoryImpl) GetPassengersBySpecificFlight(request entities.GetPassengersBySpecificFlightRequest) ([]entities.Passenger, error) {
	log.Info().Str("flight_number", request.FlightNumber).Msg("Querying passengers by specific flight")

	query := `SELECT id from flights WHERE flight_number = $1 AND departure_datetime::date = $2`
	row := r.db.QueryRow(query, request.FlightNumber, request.DepartureDateTime)

	var flight entities.Flight
	err := row.Scan(
		&flight.ID,
	)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error querying flight by flight number and departure datetime")
		return nil, err
	}

	query = `
        SELECT id, national_id, pnr_no, flight_id, payment_id, baggage_allowance, baggage_id, fare_type, seat, meal, extra_baggage, check_in, name, surname, email, phone, gender, birth_date, cip_member, vip_member, disabled, child, created_at, updated_at
        FROM passengers
        WHERE flight_id = $1
    `

	rows, err := r.db.Query(query, flight.ID)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error querying passengers by specific flight")
		return nil, err
	}
	defer rows.Close()

	var passengers []entities.Passenger
	for rows.Next() {
		var passenger entities.Passenger
		err := rows.Scan(
			&passenger.ID,
			&passenger.NationalId,
			&passenger.PnrNo,
			&passenger.FlightId,
			&passenger.PaymentId,
			&passenger.BaggageAllowance,
			&passenger.BaggageId,
			&passenger.FareType,
			&passenger.Seat,
			&passenger.Meal,
			&passenger.ExtraBaggage,
			&passenger.CheckIn,
			&passenger.Name,
			&passenger.Surname,
			&passenger.Email,
			&passenger.Phone,
			&passenger.Gender,
			&passenger.BirthDate,
			&passenger.CipMember,
			&passenger.VipMember,
			&passenger.Disabled,
			&passenger.Child,
			&passenger.CreatedAt,
			&passenger.UpdatedAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning passenger row")
			return nil, err
		}
		passengers = append(passengers, passenger)
	}

	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over passenger rows")
		return nil, err
	}

	return passengers, nil
}

func (r *PassengerRepositoryImpl) CreatePassenger(request entities.CreatePassengerRequest) error {
	log.Info().Str("national_id", request.NationalId).Msg("Creating new passenger")

	query := `
        INSERT INTO passengers (
            national_id, pnr_no, flight_id, payment_id, baggage_allowance, baggage_id, fare_type, seat, meal, extra_baggage, check_in, name, surname, email, phone, gender, birth_date, cip_member, vip_member, disabled, child
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21
        )
    `
	_, err := r.db.Exec(query,
		request.NationalId,
		request.PnrNo,
		request.FlightId,
		request.PaymentId,
		request.BaggageAllowance,
		request.BaggageId,
		request.FareType,
		request.Seat,
		request.Meal,
		request.ExtraBaggage,
		request.CheckIn,
		request.Name,
		request.Surname,
		request.Email,
		request.Phone,
		request.Gender,
		request.BirthDate,
		request.CipMember,
		request.VipMember,
		request.Disabled,
		request.Child,
	)
	if err != nil {
		log.Error().Err(err).Str("national_id", request.NationalId).Msg("Error creating passenger")
		return err
	}
	return nil
}

func (r *PassengerRepositoryImpl) GetAllPassengers() ([]entities.Passenger, error) {
	log.Info().Msg("Retrieving all passengers")

	query := `
        SELECT id, national_id, pnr_no, flight_id, payment_id, baggage_allowance, baggage_id, fare_type, seat, meal, extra_baggage, check_in, name, surname, email, phone, gender, birth_date, cip_member, vip_member, disabled, child, created_at, updated_at
        FROM passengers
    `

	rows, err := r.db.Query(query)
	if err != nil {
		log.Error().Err(err).Msg("Error querying all passengers")
		return nil, err
	}
	defer rows.Close()

	var passengers []entities.Passenger
	for rows.Next() {
		var passenger entities.Passenger
		err := rows.Scan(
			&passenger.ID,
			&passenger.NationalId,
			&passenger.PnrNo,
			&passenger.FlightId,
			&passenger.PaymentId,
			&passenger.BaggageAllowance,
			&passenger.BaggageId,
			&passenger.FareType,
			&passenger.Seat,
			&passenger.Meal,
			&passenger.ExtraBaggage,
			&passenger.CheckIn,
			&passenger.Name,
			&passenger.Surname,
			&passenger.Email,
			&passenger.Phone,
			&passenger.Gender,
			&passenger.BirthDate,
			&passenger.CipMember,
			&passenger.VipMember,
			&passenger.Disabled,
			&passenger.Child,
			&passenger.CreatedAt,
			&passenger.UpdatedAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning passenger row")
			return nil, err
		}
		passengers = append(passengers, passenger)
	}

	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over passenger rows")
		return nil, err
	}

	return passengers, nil
}

func (r *PassengerRepositoryImpl) EmployeeCheckInPassenger(request entities.EmployeeCheckInRequest) (entities.Passenger, error) {
	log.Info().
		Str("national_id", request.NationalId).
		Str("destination_airport", request.DestinationAirport).
		Msg("Employee checking in passenger")

	query := `
        SELECT p.id, p.national_id, p.pnr_no, p.flight_id, p.payment_id, 
               p.baggage_allowance, p.baggage_id, p.fare_type, p.seat, p.meal, 
               p.extra_baggage, p.check_in, p.name, p.surname, p.email, p.phone, 
               p.gender, p.birth_date, p.cip_member, p.vip_member, p.disabled, 
               p.child, p.created_at, p.updated_at
        FROM passengers p
        JOIN flights f ON p.flight_id = f.id
        WHERE p.national_id = $1 
        AND f.destination_airport = $2
        LIMIT 1`

	row := r.db.QueryRow(query, request.NationalId, request.DestinationAirport)

	var passenger entities.Passenger
	err := row.Scan(
		&passenger.ID,
		&passenger.NationalId,
		&passenger.PnrNo,
		&passenger.FlightId,
		&passenger.PaymentId,
		&passenger.BaggageAllowance,
		&passenger.BaggageId,
		&passenger.FareType,
		&passenger.Seat,
		&passenger.Meal,
		&passenger.ExtraBaggage,
		&passenger.CheckIn,
		&passenger.Name,
		&passenger.Surname,
		&passenger.Email,
		&passenger.Phone,
		&passenger.Gender,
		&passenger.BirthDate,
		&passenger.CipMember,
		&passenger.VipMember,
		&passenger.Disabled,
		&passenger.Child,
		&passenger.CreatedAt,
		&passenger.UpdatedAt,
	)
	if err != nil {
		log.Error().Err(err).
			Str("national_id", request.NationalId).
			Str("destination_airport", request.DestinationAirport).
			Msg("Error finding passenger for employee check-in")
		return entities.Passenger{}, err
	}

	// Update check-in status
	updateQuery := `UPDATE passengers SET check_in = true WHERE id = $1`
	_, err = r.db.Exec(updateQuery, passenger.ID)
	if err != nil {
		log.Error().Err(err).
			Str("passenger_id", fmt.Sprintf("%d", passenger.ID)).
			Msg("Error updating passenger check-in status")
		return entities.Passenger{}, err
	}

	return passenger, nil
}
