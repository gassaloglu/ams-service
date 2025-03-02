package mongodb

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var PLANE_LOG_PREFIX string = "plane_repository.go"

type PlaneRepositoryImpl struct {
	collection *mongo.Collection
}

func NewPlaneRepositoryImpl(client *mongo.Client, dbName, collectionName string) ports.PlaneRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &PlaneRepositoryImpl{collection: collection}
}

func (r *PlaneRepositoryImpl) GetAllPlanes() ([]entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying all planes", PLANE_LOG_PREFIX))

	ctx := context.Background()
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying all planes: %v", PLANE_LOG_PREFIX, err))
		return nil, err
	}
	defer cursor.Close(ctx)

	var planes []entities.Plane
	if err := cursor.All(ctx, &planes); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error decoding planes: %v", PLANE_LOG_PREFIX, err))
		return nil, err
	}

	return planes, nil
}

func (r *PlaneRepositoryImpl) AddPlane(request entities.AddPlaneRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Adding new plane with registration: %s", PLANE_LOG_PREFIX, request.Plane.PlaneRegistration))

	ctx := context.Background()

	// Check if a plane with the same registration already exists
	filter := bson.M{"plane_registration": request.Plane.PlaneRegistration}
	var existingPlane entities.Plane
	err := r.collection.FindOne(ctx, filter).Decode(&existingPlane)
	if err == nil {
		middlewares.LogError(fmt.Sprintf("%s - Plane with registration %s already exists", PLANE_LOG_PREFIX, request.Plane.PlaneRegistration))
		return fmt.Errorf("plane with registration %s already exists", request.Plane.PlaneRegistration)
	} else if err != mongo.ErrNoDocuments {
		middlewares.LogError(fmt.Sprintf("%s - Error checking for existing plane: %v", PLANE_LOG_PREFIX, err))
		return err
	}

	// Create a new plane document
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

	// Insert the new plane into the collection
	_, err = r.collection.InsertOne(ctx, newPlane)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error adding plane: %v", PLANE_LOG_PREFIX, err))
		return err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully added plane with registration: %s", PLANE_LOG_PREFIX, request.Plane.PlaneRegistration))
	return nil
}

func (r *PlaneRepositoryImpl) SetPlaneStatus(request entities.SetPlaneStatusRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Setting plane status for registration: %s", PLANE_LOG_PREFIX, request.PlaneRegistration))

	ctx := context.Background()

	// Update the plane's status
	filter := bson.M{"plane_registration": request.PlaneRegistration}
	update := bson.M{"$set": bson.M{"is_available": request.IsAvailable}}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error updating plane status: %v", PLANE_LOG_PREFIX, err))
		return err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully updated plane status for registration: %s", PLANE_LOG_PREFIX, request.PlaneRegistration))
	return nil
}

func (r *PlaneRepositoryImpl) GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying plane by registration: %s", PLANE_LOG_PREFIX, request.PlaneRegistration))

	ctx := context.Background()

	filter := bson.M{"plane_registration": request.PlaneRegistration}
	var plane entities.Plane
	err := r.collection.FindOne(ctx, filter).Decode(&plane)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			middlewares.LogError(fmt.Sprintf("%s - Plane with registration %s not found", PLANE_LOG_PREFIX, request.PlaneRegistration))
			return entities.Plane{}, fmt.Errorf("plane with registration %s not found", request.PlaneRegistration)
		}
		middlewares.LogError(fmt.Sprintf("%s - Error querying plane by registration: %v", PLANE_LOG_PREFIX, err))
		return entities.Plane{}, err
	}

	return plane, nil
}

func (r *PlaneRepositoryImpl) GetPlaneByFlightNumber(request entities.GetPlaneByFlightNumberRequest) (entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying plane by flight number: %s", PLANE_LOG_PREFIX, request.FlightNumber))

	ctx := context.Background()

	filter := bson.M{"flight_number": request.FlightNumber}
	var plane entities.Plane
	err := r.collection.FindOne(ctx, filter).Decode(&plane)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			middlewares.LogError(fmt.Sprintf("%s - Plane with flight number %s not found", PLANE_LOG_PREFIX, request.FlightNumber))
			return entities.Plane{}, fmt.Errorf("plane with flight number %s not found", request.FlightNumber)
		}
		middlewares.LogError(fmt.Sprintf("%s - Error querying plane by flight number: %v", PLANE_LOG_PREFIX, err))
		return entities.Plane{}, err
	}

	return plane, nil
}

func (r *PlaneRepositoryImpl) GetPlaneByLocation(request entities.GetPlaneByLocationRequest) ([]entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying planes by location: %s", PLANE_LOG_PREFIX, request.Location))

	ctx := context.Background()

	filter := bson.M{"location": request.Location}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying planes by location: %v", PLANE_LOG_PREFIX, err))
		return nil, err
	}
	defer cursor.Close(ctx)

	var planes []entities.Plane
	if err := cursor.All(ctx, &planes); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error decoding planes: %v", PLANE_LOG_PREFIX, err))
		return nil, err
	}

	return planes, nil
}
