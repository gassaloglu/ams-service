package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"database/sql"

	"github.com/rs/zerolog/log"
)

type PlaneRepositoryImpl struct {
	db *sql.DB
}

func NewPlaneRepositoryImpl(db *sql.DB) secondary.PlaneRepository {
	return &PlaneRepositoryImpl{db: db}
}

func (r *PlaneRepositoryImpl) GetAllPlanes() ([]entities.Plane, error) {
	log.Info().Msg("Getting all planes")

	query := "SELECT * FROM planes"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var planes []entities.Plane
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var plane entities.Plane
		err := rows.Scan(&plane.PlaneRegistration, &plane.PlaneType, &plane.Location, &plane.TotalPassengers, &plane.MaxPassengers, &plane.EconomyPassengers, &plane.BusinessPassengers, &plane.FirstClassPassengers, &plane.FlightNumber, &plane.IsAvailable)
		if err != nil {
			return planes, err
		}
		planes = append(planes, plane)
	}
	if err = rows.Err(); err != nil {
		return planes, err
	}
	return planes, nil
}

func (r *PlaneRepositoryImpl) AddPlane(request entities.AddPlaneRequest) error {
	// Check whether there are any planes that has same registration number
	getByRegistrationInput := entities.GetPlaneByRegistrationRequest{PlaneRegistration: request.Plane.PlaneRegistration}
	existedPlane, err := r.GetPlaneByRegistration(getByRegistrationInput)
	var plane entities.Plane
	if existedPlane != plane {
		log.Info().Str("registration", request.Plane.PlaneRegistration).Msg("Plane with the registration already exists")
		return err
	}
	log.Info().Str("registration", request.Plane.PlaneRegistration).Msg("Adding new plane")

	query := `
			INSERT INTO planes (
				plane_registration, plane_type, location, total_passengers, max_passengers,
				economy_passengers, business_passengers, first_class_passengers, flight_number, is_available
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		`
	_, err = r.db.Exec(
		query,
		request.Plane.PlaneRegistration,
		request.Plane.PlaneType,
		request.Plane.Location,
		request.Plane.TotalPassengers,
		request.Plane.MaxPassengers,
		request.Plane.EconomyPassengers,
		request.Plane.BusinessPassengers,
		request.Plane.FirstClassPassengers,
		request.Plane.FlightNumber,
		request.Plane.IsAvailable,
	)
	if err != nil {
		log.Error().Err(err).Str("registration", request.Plane.PlaneRegistration).Msg("Error adding plane")
		return err
	}
	return nil
}

func (r *PlaneRepositoryImpl) SetPlaneStatus(request entities.SetPlaneStatusRequest) error {
	log.Info().Str("registration", request.PlaneRegistration).Bool("is_available", request.IsAvailable).Msg("Setting plane status")

	query := "UPDATE planes SET is_available = $1 WHERE plane_registration = $2"
	_, err := r.db.Exec(query, request.IsAvailable, request.PlaneRegistration)
	if err != nil {
		log.Error().Err(err).Str("registration", request.PlaneRegistration).Msg("Error updating plane status")
		return err
	}
	return nil
}

func (r *PlaneRepositoryImpl) GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error) {
	log.Info().Str("registration", request.PlaneRegistration).Msg("Querying plane by registration")

	query := "SELECT * FROM planes WHERE plane_registration = $1"
	row := r.db.QueryRow(query, request.PlaneRegistration)

	var plane entities.Plane

	err := row.Scan(&plane.PlaneRegistration, &plane.PlaneType, &plane.Location, &plane.TotalPassengers, &plane.MaxPassengers, &plane.EconomyPassengers, &plane.BusinessPassengers, &plane.FirstClassPassengers, &plane.FlightNumber, &plane.IsAvailable)
	if err != nil {
		log.Error().Err(err).Str("registration", request.PlaneRegistration).Msg("Error querying plane by registration")
		return entities.Plane{}, err
	}
	return plane, nil
}

func (r *PlaneRepositoryImpl) GetPlaneByFlightNumber(request entities.GetPlaneByFlightNumberRequest) (entities.Plane, error) {
	log.Info().Str("flight_number", request.FlightNumber).Msg("Querying plane by flight number")

	query := "SELECT * FROM planes WHERE flight_number = $1"
	row := r.db.QueryRow(query, request.FlightNumber)

	var plane entities.Plane

	err := row.Scan(&plane.PlaneRegistration, &plane.PlaneType, &plane.Location, &plane.TotalPassengers, &plane.MaxPassengers, &plane.EconomyPassengers, &plane.BusinessPassengers, &plane.FirstClassPassengers, &plane.FlightNumber, &plane.IsAvailable)
	if err != nil {
		log.Error().Err(err).Str("flight_number", request.FlightNumber).Msg("Error querying plane by flight number")
		return entities.Plane{}, err
	}
	return plane, nil
}

func (r *PlaneRepositoryImpl) GetPlaneByLocation(request entities.GetPlaneByLocationRequest) ([]entities.Plane, error) {
	log.Info().Str("location", request.Location).Msg("Querying planes by location")

	query := "SELECT * FROM planes WHERE location = $1"
	rows, err := r.db.Query(query, request.Location)
	if err != nil {
		log.Error().Err(err).Str("location", request.Location).Msg("Error querying planes by location")
		return nil, err
	}
	defer rows.Close()

	var planes []entities.Plane
	for rows.Next() {
		var plane entities.Plane
		err := rows.Scan(&plane.PlaneRegistration, &plane.PlaneType, &plane.Location, &plane.TotalPassengers, &plane.MaxPassengers, &plane.EconomyPassengers, &plane.BusinessPassengers, &plane.FirstClassPassengers, &plane.FlightNumber, &plane.IsAvailable)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning plane row")
			return planes, err
		}
		planes = append(planes, plane)
	}
	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over plane rows")
		return planes, err
	}
	return planes, nil
}
