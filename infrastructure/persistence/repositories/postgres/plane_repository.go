package postgres

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"database/sql"
	"fmt"
)

var PLANE_LOG_PREFIX string = "passenger_repository.go"

type PlaneRepositoryImpl struct {
	db *sql.DB
}

func NewPlaneRepositoryImpl(db *sql.DB) ports.PlaneRepository {
	return &PlaneRepositoryImpl{db: db}
}

func (r *PlaneRepositoryImpl) GetAllPlanes() ([]entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("Getting all planes"))

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
		middlewares.LogInfo(fmt.Sprintf("%s - Plane with the registration: %s is already existed", PLANE_LOG_PREFIX, request.Plane.PlaneRegistration))
		return err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Adding new plane with registration: %s", PLANE_LOG_PREFIX, request.Plane.PlaneRegistration))

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
		middlewares.LogError(fmt.Sprintf("%s - Error adding plane: %v", PLANE_LOG_PREFIX, err))
		return err
	}
	return nil
}

func (r *PlaneRepositoryImpl) SetPlaneStatus(request entities.SetPlaneStatusRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Setting plane status for registration: %s", PLANE_LOG_PREFIX, request.PlaneRegistration))

	query := "UPDATE planes SET is_available = $1 WHERE plane_registration = $2"
	_, err := r.db.Exec(query, request.IsAvailable, request.PlaneRegistration)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error updating plane status: %v", PLANE_LOG_PREFIX, err))
		return err
	}
	return nil
}

func (r *PlaneRepositoryImpl) GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying plane by registrtion: %s", PLANE_LOG_PREFIX, request.PlaneRegistration))

	query := "SELECT * FROM planes WHERE regstration = $1"
	row := r.db.QueryRow(query, request.PlaneRegistration)

	var plane entities.Plane

	err := row.Scan(&plane.PlaneRegistration, &plane.PlaneType, &plane.Location, &plane.TotalPassengers, &plane.MaxPassengers, &plane.EconomyPassengers, &plane.BusinessPassengers, &plane.FirstClassPassengers, &plane.FlightNumber, &plane.IsAvailable)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying plane by registration: %v", PLANE_LOG_PREFIX, err))
		return entities.Plane{}, err
	}
	return plane, nil
}

func (r *PlaneRepositoryImpl) GetPlaneByFlightNumber(request entities.GetPlaneByFlightNumberRequest) (entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying plane by flight number: %s", PLANE_LOG_PREFIX, request.FlightNumber))

	query := "SELECT * FROM planes WHERE flight_number = $1"
	row := r.db.QueryRow(query, request.FlightNumber)

	var plane entities.Plane

	err := row.Scan(&plane.PlaneRegistration, &plane.PlaneType, &plane.Location, &plane.TotalPassengers, &plane.MaxPassengers, &plane.EconomyPassengers, &plane.BusinessPassengers, &plane.FirstClassPassengers, &plane.FlightNumber, &plane.IsAvailable)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying plane by flight number: %v", PLANE_LOG_PREFIX, err))
		return entities.Plane{}, err
	}
	return plane, nil
}

func (r *PlaneRepositoryImpl) GetPlaneByLocation(request entities.GetPlaneByLocationRequest) ([]entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying planes by location: %s", PLANE_LOG_PREFIX, request.Location))

	query := "SELECT * FROM planes WHERE location = $1"
	rows, err := r.db.Query(query, request.Location)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying planes by location: %v", PLANE_LOG_PREFIX, err))
		return nil, err
	}
	defer rows.Close()

	var planes []entities.Plane
	for rows.Next() {
		var plane entities.Plane
		err := rows.Scan(&plane.PlaneRegistration, &plane.PlaneType, &plane.Location, &plane.TotalPassengers, &plane.MaxPassengers, &plane.EconomyPassengers, &plane.BusinessPassengers, &plane.FirstClassPassengers, &plane.FlightNumber, &plane.IsAvailable)
		if err != nil {
			middlewares.LogError(fmt.Sprintf("%s - Error scanning plane row: %v", PLANE_LOG_PREFIX, err))
			return planes, err
		}
		planes = append(planes, plane)
	}
	if err = rows.Err(); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error iterating over plane rows: %v", PLANE_LOG_PREFIX, err))
		return planes, err
	}
	return planes, nil
}
