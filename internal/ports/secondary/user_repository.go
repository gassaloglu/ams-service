package secondary

import (
	"ams-service/internal/core/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	CreateAll(user []entities.User) error
	FindUserByEmail(email string) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
}
