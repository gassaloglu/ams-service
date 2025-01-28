package mongodb

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var USER_LOG_PREFIX string = "user_repository.go"

type UserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepositoryImpl(client *mongo.Client, dbName, collectionName string) ports.UserRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &UserRepositoryImpl{collection: collection}
}

func (r *UserRepositoryImpl) RegisterUser(user entities.User) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Registering user: %v", USER_LOG_PREFIX, user))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering user: %v", USER_LOG_PREFIX, err))
		return err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully registered user: %v", USER_LOG_PREFIX, user))
	return nil
}
