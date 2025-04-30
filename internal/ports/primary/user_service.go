package primary

import (
	"ams-service/internal/core/entities"
)

type UserService interface {
	Register(user *entities.UserRegisterRequest) (string, error)
	Login(email, password string) (string, error)
}
