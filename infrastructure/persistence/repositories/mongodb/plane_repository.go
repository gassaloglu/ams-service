package mongodb

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var PLANE_LOG_PREFIX string = "passenger_repository.go"

type PlaneRepositoryImpl struct {
	collection *mongo.Collection
}

func NewPlaneRepositoryImpl(client *mongo.Client, dbName, collectionName string) ports.PlaneRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &PlaneRepositoryImpl{collection: collection}
}

func (r *PlaneRepositoryImpl) GetAllPlanes() ([]entities.Plane, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying all planes", PASSENGER_LOG_PREFIX))

	// Will be updated

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	// var passenger []entities.Plane
	// err := r.collection.FindOne(ctx, bson.M{}).Decode(&passenger)
	// if err != nil {
	// 	middlewares.LogError(fmt.Sprintf("%s - Error querying passenger by ID: %v", PASSENGER_LOG_PREFIX, err))
	// 	return []entities.Plane{}, err
	// }

	//return passenger, nil
	return nil, nil
}

func (r *PlaneRepositoryImpl) AddPlane(request entities.AddPlaneRequest) error {
	// Will be added
	return nil
}

func (r *PlaneRepositoryImpl) SetPlaneStatus(request entities.SetPlaneStatusRequest) error {
	// Will be added
	return nil
}

func (s *PlaneRepositoryImpl) GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error) {
	// Will be added
	return entities.Plane{}, nil
}

func (s *PlaneRepositoryImpl) GetPlaneByFlightNumber(request entities.GetPlaneByFlightNumberRequest) (entities.Plane, error) {
	// Will be added
	return entities.Plane{}, nil
}

func (s *PlaneRepositoryImpl) GetPlaneByLocation(request entities.GetPlaneByLocationRequest) ([]entities.Plane, error) {
	// Will be added
	return []entities.Plane{}, nil
}
