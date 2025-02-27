package ports

import "ams-service/core/entities"

type EmployeeRepository interface {
	RegisterEmployee(request entities.RegisterEmployeeRequest) error
	GetEmployeeByNationalID(entities.GetEmployeeByNationalIDRequest) (entities.Employee, error)
}
