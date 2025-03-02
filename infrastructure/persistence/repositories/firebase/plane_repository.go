package firebase

/*
import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"context"
	"fmt"

	"firebase.google.com/go/v4/db"
)

var PLANE_LOG_PREFIX string = "plane_repository.go"

type PlaneRepositoryImpl struct {
	client *db.Client
}

func NewPlaneRepositoryImpl(client *db.Client) ports.PlaneRepository {
	return &PlaneRepositoryImpl{client: client}
}

func (r *PlaneRepositoryImpl) GetAllPlanes() ([]entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying all planes", PLANE_LOG_PREFIX))

	ctx := context.Background()
	ref := r.client.NewRef("planes")

	var planes map[string]entities.Plane
	err := ref.Get(ctx, &planes)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying all planes: %v", PLANE_LOG_PREFIX, err))
		return []entities.Plane{}, err
	}

	// Convert map to slice
	var planeList []entities.Plane
	for _, plane := range planes {
		planeList = append(planeList, plane)
	}

	return planeList, nil
}

func (r *PlaneRepositoryImpl) AddPlane(request entities.AddPlaneRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Adding new plane with registration: %s", PLANE_LOG_PREFIX, request.Plane.PlaneRegistration))

	ctx := context.Background()
	ref := r.client.NewRef("planes")

	// Check if a plane with the same registration already exists
	existingPlane, _ := r.GetPlaneByRegistration(entities.GetPlaneByRegistrationRequest{
		PlaneRegistration: request.Plane.PlaneRegistration,
	})
	var plane entities.Plane
	if existingPlane != plane {
		middlewares.LogError(fmt.Sprintf("%s - Plane with registration %s already exists", PLANE_LOG_PREFIX, request.Plane.PlaneRegistration))
		return fmt.Errorf("plane with registration %s already exists", request.Plane.PlaneRegistration)
	}

	// Create a new plane entry
	newPlane := entities.Plane{
		PlaneRegistration:    request.Plane.PlaneRegistration,
		PlaneType:            request.Plane.PlaneType,
		Location:             request.Plane.Location,
		TotalPassengers:      request.Plane.TotalPassengers,
		MaxPassengers:        request.Plane.MaxPassengers,
		EconomyPassengers:    request.Plane.EconomyPassengers,
		BusinessPassengers:   request.Plane.BusinessPassengers,
		FirstClassPassengers: request.Plane.FirstClassPassengers,
		FlightNumber:         request.Plane.FlightNumber,
		IsAvailable:          request.Plane.IsAvailable,
	}

	// Add the new plane to Firebase
	planeRef := ref.Child(request.Plane.PlaneRegistration)
	if err := planeRef.Set(ctx, newPlane); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error adding plane: %v", PLANE_LOG_PREFIX, err))
		return err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully added plane with registration: %s", PLANE_LOG_PREFIX, request.Plane.PlaneRegistration))
	return nil
}

func (r *PlaneRepositoryImpl) SetPlaneStatus(request entities.SetPlaneStatusRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Setting plane status for registration: %s", PLANE_LOG_PREFIX, request.PlaneRegistration))

	ctx := context.Background()
	ref := r.client.NewRef("planes")

	// Update the plane's status
	planeRef := ref.Child(request.PlaneRegistration)
	if err := planeRef.Update(ctx, map[string]interface{}{
		"is_available": request.IsAvailable,
	}); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error updating plane status: %v", PLANE_LOG_PREFIX, err))
		return err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully updated plane status for registration: %s", PLANE_LOG_PREFIX, request.PlaneRegistration))
	return nil
}

func (r *PlaneRepositoryImpl) GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying plane by registration: %s", PLANE_LOG_PREFIX, request.PlaneRegistration))

	ctx := context.Background()
	ref := r.client.NewRef("planes")

	var plane entities.Plane
	planeRef := ref.Child(request.PlaneRegistration)
	if err := planeRef.Get(ctx, &plane); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying plane by registration: %v", PLANE_LOG_PREFIX, err))
		return entities.Plane{}, err
	}

	return plane, nil
}

func (r *PlaneRepositoryImpl) GetPlaneByFlightNumber(request entities.GetPlaneByFlightNumberRequest) (entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying plane by flight number: %s", PLANE_LOG_PREFIX, request.FlightNumber))

	ctx := context.Background()
	ref := r.client.NewRef("planes")

	var planes map[string]entities.Plane
	if err := ref.Get(ctx, &planes); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying planes: %v", PLANE_LOG_PREFIX, err))
		return entities.Plane{}, err
	}

	// Find the plane with the matching flight number
	for _, plane := range planes {
		if plane.FlightNumber == request.FlightNumber {
			return plane, nil
		}
	}

	middlewares.LogError(fmt.Sprintf("%s - Plane with flight number %s not found", PLANE_LOG_PREFIX, request.FlightNumber))
	return entities.Plane{}, fmt.Errorf("plane with flight number %s not found", request.FlightNumber)
}

func (r *PlaneRepositoryImpl) GetPlaneByLocation(request entities.GetPlaneByLocationRequest) ([]entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying planes by location: %s", PLANE_LOG_PREFIX, request.Location))

	ctx := context.Background()
	ref := r.client.NewRef("planes")

	var planes map[string]entities.Plane
	if err := ref.Get(ctx, &planes); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying planes: %v", PLANE_LOG_PREFIX, err))
		return []entities.Plane{}, err
	}

	// Filter planes by location
	var filteredPlanes []entities.Plane
	for _, plane := range planes {
		if plane.Location == request.Location {
			filteredPlanes = append(filteredPlanes, plane)
		}
	}

	return filteredPlanes, nil
}
*/
