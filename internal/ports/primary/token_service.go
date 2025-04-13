package primary

import "ams-service/internal/core/entities"

type Token string

type TokenService interface {
	CreateUserToken(user *entities.User) (string, error)
	CreateEmployeeToken(employee *entities.Employee) (string, error)
	ValidateToken(token string) error
	ValidateUserToken(token string) error
	ValidateEmployeeToken(token string) error
	ValidateRole(token string, allowedRoles []string) error
}
