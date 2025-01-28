package services

import (
	"ams-service/core/entities"
	"ams-service/middlewares"
	"fmt"
)

var USER_LOG_PREFIX string = "user_service.go"

type UserRepository interface {
	RegisterUser(user entities.User) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user entities.User) error {
	err := s.repo.RegisterUser(user)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering user: %v", USER_LOG_PREFIX, err))
		return err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully registered user: %v", USER_LOG_PREFIX, user))
	return nil
}
