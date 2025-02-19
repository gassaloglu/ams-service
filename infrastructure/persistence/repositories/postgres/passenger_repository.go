package postgres

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"database/sql"
	"fmt"
)

var PASSENGER_LOG_PREFIX string = "passenger_repository.go"

type PassengerRepositoryImpl struct {
	db *sql.DB
}

func NewPassengerRepositoryImpl(db *sql.DB) ports.PassengerRepository {
	return &PassengerRepositoryImpl{db: db}
}

func (r *PassengerRepositoryImpl) GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying passenger by ID: %s", PASSENGER_LOG_PREFIX, request.NationalId))
	query := `SELECT * FROM passengers WHERE id = $1`
	row := r.db.QueryRow(query, request.NationalId)

	var passenger entities.Passenger
	err := row.Scan(&passenger.NationalId, &passenger.PnrNo, &passenger.BaggageAllowance, &passenger.BaggageId, &passenger.FareType, &passenger.Seat, &passenger.Meal, &passenger.ExtraBaggage, &passenger.CheckIn, &passenger.Name, &passenger.Surname, &passenger.Email, &passenger.Phone, &passenger.Gender, &passenger.BirthDate, &passenger.CipMember, &passenger.VipMember, &passenger.Disabled, &passenger.Child)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying passenger by ID: %v", PASSENGER_LOG_PREFIX, err))
		return entities.Passenger{}, err
	}

	return passenger, nil
}

func (r *PassengerRepositoryImpl) GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying passenger by PNR: %s and surname: %s", PASSENGER_LOG_PREFIX, request.PNR, request.Surname))
	query := `SELECT * FROM passengers WHERE pnr_no = $1 AND surname = $2`
	row := r.db.QueryRow(query, request.PNR, request.Surname)

	var passenger entities.Passenger
	err := row.Scan(&passenger.NationalId, &passenger.PnrNo, &passenger.BaggageAllowance, &passenger.BaggageId, &passenger.FareType, &passenger.Seat, &passenger.Meal, &passenger.ExtraBaggage, &passenger.CheckIn, &passenger.Name, &passenger.Surname, &passenger.Email, &passenger.Phone, &passenger.Gender, &passenger.BirthDate, &passenger.CipMember, &passenger.VipMember, &passenger.Disabled, &passenger.Child)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying passenger by PNR: %v", PASSENGER_LOG_PREFIX, err))
		return entities.Passenger{}, err
	}

	return passenger, nil
}

func (r *PassengerRepositoryImpl) OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Checking in passenger with PNR: %s and surname: %s", PASSENGER_LOG_PREFIX, request.PNR, request.Surname))
	query := `UPDATE passengers SET check_in = true WHERE pnr_no = $1 AND surname = $2`
	_, err := r.db.Exec(query, request.PNR, request.Surname)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error checking in passenger: %v", PASSENGER_LOG_PREFIX, err))
		return err
	}

	return nil
}
