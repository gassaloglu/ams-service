package primary

import (
	"ams-service/internal/core/entities"
)

type UserService interface {
	Register(user *entities.RegisterUserRequest) (string, error)
	RegisterAll(user []entities.RegisterUserRequest) error
	Login(email, password string) (string, error)
	GetAllUsers() ([]entities.User, error)
}
