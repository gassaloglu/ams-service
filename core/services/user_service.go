package services

import (
	"ams-service/core/entities"
	"ams-service/middlewares"
	"ams-service/utils"
	"fmt"
)

var USER_LOG_PREFIX string = "user_service.go"

type UserRepository interface {
	RegisterUser(user entities.User) error
	LoginUser(username, password string) (*entities.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user entities.User) error {
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error generating salt: %v", USER_LOG_PREFIX, err))
		return err
	}

	hashedPassword, err := utils.HashPassword(user.PasswordHash, salt)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error hashing password: %v", USER_LOG_PREFIX, err))
		return err
	}

	user.PasswordHash = hashedPassword
	user.Salt = salt

	err = s.repo.RegisterUser(user)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering user: %v", USER_LOG_PREFIX, err))
		return err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully registered user: %v", USER_LOG_PREFIX, user))
	return nil
}

func (s *UserService) LoginUser(username, password string) (*entities.User, error) {
	user, err := s.repo.LoginUser(username, password)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error logging in user: %v", USER_LOG_PREFIX, err))
		return nil, err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully logged in user: %s", USER_LOG_PREFIX, username))
	return user, nil
}
