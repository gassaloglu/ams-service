package primary

import (
	"ams-service/internal/core/entities"
)

type UserService interface {
	RegisterUser(user entities.User) error
	LoginUser(username, password string) (*entities.User, string, error)
}
