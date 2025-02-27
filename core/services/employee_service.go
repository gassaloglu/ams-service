package services

import (
	"ams-service/core/entities"
	"ams-service/middlewares"
	"fmt"
)

var EMPLOYEE_LOG_PREFIX string = "employee_service.go"

type EmployeeRepository interface {
	RegisterEmployee(request entities.RegisterEmployeeRequest) error
	GetEmployeeByNationalID(entities.GetEmployeeByNationalIDRequest) (entities.Employee, error)
	GetEmployeeByID(entities.GetEmployeeByIDRequest) (entities.Employee, error)
}

type EmployeeService struct {
	repo EmployeeRepository
}

func NewEmployeeService(repo EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) RegisterEmployee(request entities.RegisterEmployeeRequest) error {
	err := s.repo.RegisterEmployee(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully registered employee", EMPLOYEE_LOG_PREFIX))
	return nil
}

func (s *EmployeeService) GetEmployeeByNationalID(request entities.GetEmployeeByNationalIDRequest) (entities.Employee, error) {
	employee, err := s.repo.GetEmployeeByNationalID(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting employee by national ID %s: %v", EMPLOYEE_LOG_PREFIX, request.EmployeeID, err))
		return entities.Employee{}, err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved employee by national ID %s", EMPLOYEE_LOG_PREFIX, request.EmployeeID))
	return employee, nil
}

func (s *EmployeeService) GetEmployeeByID(request entities.GetEmployeeByIDRequest) (entities.Employee, error) {
	employee, err := s.repo.GetEmployeeByID(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting employee by ID %d: %v", EMPLOYEE_LOG_PREFIX, request.ID, err))
		return entities.Employee{}, err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully retrieved employee by ID %d", EMPLOYEE_LOG_PREFIX, request.ID))
	return employee, nil
}
