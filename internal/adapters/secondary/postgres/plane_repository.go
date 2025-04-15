package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"database/sql"
	"time"

	"github.com/rs/zerolog/log"
)

type PlaneRepositoryImpl struct {
	db *sql.DB
}

func NewPlaneRepositoryImpl(db *sql.DB) secondary.PlaneRepository {
	return &PlaneRepositoryImpl{db: db}
}

func (r *PlaneRepositoryImpl) GetAllPlanes(req entities.GetAllPlanesRequest) ([]entities.Plane, error) {
	log.Info().Msg("Getting all planes")

	query := `
		SELECT * FROM planes
	`

	// FIXME: Add filtering by capacity and status
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	var planes []entities.Plane
	for rows.Next() {
		var plane entities.Plane
		err := rows.Scan(&plane.Registration, &plane.Model, &plane.Manufacturer, &plane.Capacity, &plane.Status, &plane.CreatedAt, &plane.UpdatedAt)
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
	for _, plane := range request {
		query := `
				INSERT INTO planes (
					registration, model, manufacturer, capacity, status, created_at, updated_at
				) VALUES ($1, $2, $3, $4, $5, $6, $7)
			`
		_, err := r.db.Exec(
			query,
			plane.Registration,
			plane.Model,
			plane.Manufacturer,
			plane.Capacity,
			plane.Status,
			time.Now(),
			time.Now(),
		)

		if err != nil {
			log.Error().Err(err).Str("registration", plane.Registration).Msg("Error adding plane")
			return err
		}
	}

	return nil
}

func (r *PlaneRepositoryImpl) SetPlaneStatus(request entities.SetPlaneStatusRequest) error {
	log.Info().Str("registration", request.PlaneRegistration).Str("status", request.Status).Msg("Setting plane status")

	query := "UPDATE planes SET status = $1 WHERE registration = $2"
	_, err := r.db.Exec(query, request.Status, request.PlaneRegistration)
	if err != nil {
		log.Error().Err(err).Str("registration", request.PlaneRegistration).Msg("Error updating plane status")
		return err
	}
	return nil
}

func (r *PlaneRepositoryImpl) GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error) {
	log.Info().Str("registration", request.Registration).Msg("Querying plane by registration")

	query := "SELECT * FROM planes WHERE registration = $1"
	row := r.db.QueryRow(query, request.Registration)

	var plane entities.Plane

	err := row.Scan(&plane.Registration, &plane.Model, &plane.Manufacturer, &plane.Capacity, &plane.Status, &plane.CreatedAt, &plane.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Str("registration", request.Registration).Msg("Error querying plane by registration")
		return entities.Plane{}, err
	}
	return plane, nil
}
