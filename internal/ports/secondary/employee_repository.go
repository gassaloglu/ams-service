package secondary

import (
	"ams-service/internal/core/entities"
)

type EmployeeRepository interface {
	FindAll() ([]entities.Employee, error)
	FindByNationalId(nationalId string) (*entities.Employee, error)
	Create(employee *entities.Employee) (*entities.Employee, error)
	CreateAll(employee []entities.Employee) error
}
