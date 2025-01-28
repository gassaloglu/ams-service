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

func (r *PassengerRepositoryImpl) GetPassengerByID(passengerID string) (entities.Passenger, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying passenger by ID: %s", PASSENGER_LOG_PREFIX, passengerID))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var passenger entities.Passenger
	err := r.collection.FindOne(ctx, bson.M{"id": passengerID}).Decode(&passenger)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying passenger by ID: %v", PASSENGER_LOG_PREFIX, err))
		return entities.Passenger{}, err
	}

	return passenger, nil
}

func (r *PassengerRepositoryImpl) OnlineCheckInPassenger(pnr, surname string) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Checking in passenger with PNR: %s and surname: %s", PASSENGER_LOG_PREFIX, pnr, surname))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"pnr_no": pnr, "surname": surname}
	update := bson.M{"$set": bson.M{"check_in": true}}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error checking in passenger: %v", PASSENGER_LOG_PREFIX, err))
		return err
	}

	return nil
}
