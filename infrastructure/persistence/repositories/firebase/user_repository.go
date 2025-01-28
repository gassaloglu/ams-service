package firebase

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"context"
	"fmt"

	"firebase.google.com/go/v4/db"
)

var USER_LOG_PREFIX string = "user_repository.go"

type UserRepositoryImpl struct {
	client *db.Client
}

func NewUserRepositoryImpl(client *db.Client) ports.UserRepository {
	return &UserRepositoryImpl{client: client}
}

func (r *UserRepositoryImpl) RegisterUser(user entities.User) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Registering user: %v", USER_LOG_PREFIX, user))

	ctx := context.Background()
	ref := r.client.NewRef("users")

	_, err := ref.Push(ctx, user)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering user: %v", USER_LOG_PREFIX, err))
		return err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully registered user: %v", USER_LOG_PREFIX, user))
	return nil
}
