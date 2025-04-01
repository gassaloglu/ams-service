package primary

import "ams-service/internal/core/entities"

type EmployeeService interface {
	GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error)
	RegisterEmployee(request entities.RegisterEmployeeRequest) error
	LoginEmployee(employeeID, password string) (*entities.Employee, string, error)
}
