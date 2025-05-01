package primary

import "ams-service/internal/core/entities"

type EmployeeService interface {
	FindAll() ([]entities.Employee, error)
	RegisterAll(request []entities.RegisterEmployeeRequest) error
	Register(request *entities.RegisterEmployeeRequest) (string, error)
	Login(request *entities.LoginEmployeeRequest) (string, error)
}
