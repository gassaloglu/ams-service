package ports

import "ams-service/core/entities"

type UserRepository interface {
	RegisterUser(user entities.User) error
}
