package ports

import "ams-service/core/entities"

type EmployeeRepository interface {
	GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error)
	RegisterEmployee(request entities.RegisterEmployeeRequest) error
	LoginEmployee(employeeID, password string) (*entities.Employee, error)
}
