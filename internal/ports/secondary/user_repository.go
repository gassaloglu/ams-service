package secondary

import (
	"ams-service/internal/core/entities"
)

type UserRepository interface {
	RegisterUser(user entities.User) error
	LoginUser(username, password string) (*entities.User, error)
}
