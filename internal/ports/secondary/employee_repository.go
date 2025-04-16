package secondary

import (
	"ams-service/internal/core/entities"
)

type EmployeeRepository interface {
	GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error)
	RegisterEmployee(request entities.Employee) error
	LoginEmployee(employeeID, password string) (*entities.Employee, error)
}
