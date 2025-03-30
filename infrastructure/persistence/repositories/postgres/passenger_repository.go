package postgres

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"database/sql"

	"github.com/rs/zerolog/log"
)

var PASSENGER_LOG_PREFIX string = "passenger_repository.go"

type PassengerRepositoryImpl struct {
	db *sql.DB
}

func NewPassengerRepositoryImpl(db *sql.DB) ports.PassengerRepository {
	return &PassengerRepositoryImpl{db: db}
}

func (r *PassengerRepositoryImpl) GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error) {
	log.Info().Str("national_id", request.NationalId).Msg("Querying passenger by ID")
	query := `SELECT * FROM passengers WHERE id = $1`
	row := r.db.QueryRow(query, request.NationalId)

	var passenger entities.Passenger
	err := row.Scan(&passenger.NationalId, &passenger.PnrNo, &passenger.BaggageAllowance, &passenger.BaggageId, &passenger.FareType, &passenger.Seat, &passenger.Meal, &passenger.ExtraBaggage, &passenger.CheckIn, &passenger.Name, &passenger.Surname, &passenger.Email, &passenger.Phone, &passenger.Gender, &passenger.BirthDate, &passenger.CipMember, &passenger.VipMember, &passenger.Disabled, &passenger.Child)
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
	err := row.Scan(&passenger.NationalId, &passenger.PnrNo, &passenger.BaggageAllowance, &passenger.BaggageId, &passenger.FareType, &passenger.Seat, &passenger.Meal, &passenger.ExtraBaggage, &passenger.CheckIn, &passenger.Name, &passenger.Surname, &passenger.Email, &passenger.Phone, &passenger.Gender, &passenger.BirthDate, &passenger.CipMember, &passenger.VipMember, &passenger.Disabled, &passenger.Child)
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
