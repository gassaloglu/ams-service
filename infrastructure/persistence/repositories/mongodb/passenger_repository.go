package mongodb

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var PASSENGER_LOG_PREFIX string = "passenger_repository.go"

type PassengerRepositoryImpl struct {
	collection *mongo.Collection
}

func NewPassengerRepositoryImpl(client *mongo.Client, dbName, collectionName string) ports.PassengerRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &PassengerRepositoryImpl{collection: collection}
}

func (r *PassengerRepositoryImpl) GetPassengerByID(request entities.GetPassengerByIdRequest) (entities.Passenger, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying passenger by ID: %s", PASSENGER_LOG_PREFIX, request.NationalId))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var passenger entities.Passenger
	err := r.collection.FindOne(ctx, bson.M{"id": request.NationalId}).Decode(&passenger)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying passenger by ID: %v", PASSENGER_LOG_PREFIX, err))
		return entities.Passenger{}, err
	}

	return passenger, nil
}

func (r *PassengerRepositoryImpl) GetPassengerByPNR(request entities.GetPassengerByPnrRequest) (entities.Passenger, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying passenger by PNR: %s and surname: %s", PASSENGER_LOG_PREFIX, request.PNR, request.Surname))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"pnr_no": request.PNR, "surname": request.Surname}

	var passenger entities.Passenger
	err := r.collection.FindOne(ctx, filter).Decode(&passenger)
	if err != nil {
		middlewares.LogInfo(fmt.Sprintf("%s - Error querying passenger by PNR: %v", PASSENGER_LOG_PREFIX, err))
		return entities.Passenger{}, err
	}

	return passenger, nil
}

func (r *PassengerRepositoryImpl) OnlineCheckInPassenger(request entities.OnlineCheckInRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Checking in passenger with PNR: %s and surname: %s", PASSENGER_LOG_PREFIX, request.PNR, request.Surname))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"pnr_no": request.PNR, "surname": request.Surname}
	update := bson.M{"$set": bson.M{"check_in": true}}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error checking in passenger: %v", PASSENGER_LOG_PREFIX, err))
		return err
	}

	return nil
}
