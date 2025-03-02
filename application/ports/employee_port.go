package ports

import "ams-service/core/entities"

type EmployeeRepository interface {
	GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error)
}
