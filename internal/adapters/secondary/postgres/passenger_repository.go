package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"database/sql"

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
	query := `SELECT * FROM passengers WHERE national_id = $1`
	row := r.db.QueryRow(query, request.NationalId)

	var passenger entities.Passenger
	err := row.Scan(
		&passenger.ID,
		&passenger.NationalId,
		&passenger.PnrNo,
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
	query := `SELECT * FROM passengers WHERE pnr_no = $1 AND surname = $2`
	row := r.db.QueryRow(query, request.PNR, request.Surname)

	var passenger entities.Passenger
	err := row.Scan(
		&passenger.ID,
		&passenger.NationalId,
		&passenger.PnrNo,
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

	query := `
        SELECT national_id, pnr_no, baggage_allowance, baggage_id, fare_type, seat, meal, extra_baggage, check_in, name, surname, email, phone, gender, birth_date, cip_member, vip_member, disabled, child
        FROM passengers
        WHERE flight_number = $1
    `

	rows, err := r.db.Query(query, request.FlightNumber)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error querying passengers by specific flight")
		return nil, err
	}
	defer rows.Close()

	var passengers []entities.Passenger
	for rows.Next() {
		var passenger entities.Passenger
		err := rows.Scan(
			&passenger.NationalId,
			&passenger.PnrNo,
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
            national_id, pnr_no, baggage_allowance, baggage_id, fare_type, seat, meal, extra_baggage, check_in, name, surname, email, phone, gender, birth_date, cip_member, vip_member, disabled, child
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
        )
    `
	_, err := r.db.Exec(query,
		request.NationalId,
		request.PnrNo,
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
        SELECT national_id, pnr_no, baggage_allowance, baggage_id, fare_type, seat, meal, extra_baggage, check_in, name, surname, email, phone, gender, birth_date, cip_member, vip_member, disabled, child
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
			&passenger.NationalId,
			&passenger.PnrNo,
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
