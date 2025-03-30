package services

import (
	"ams-service/config"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"ams-service/utils"
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var EMPLOYEE_LOG_PREFIX string = "employee_service.go"
var EMPLOYEE_TOKEN_EXPIRY_DURATION time.Duration = time.Hour * 72

type EmployeeRepository interface {
	GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error)
	RegisterEmployee(request entities.RegisterEmployeeRequest) error
	LoginEmployee(employeeID, password string) (*entities.Employee, error)
}
type EmployeeService struct {
	repo EmployeeRepository
}

func NewEmployeeService(repo EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error) {
	employee, err := s.repo.GetEmployeeByID(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting employee by ID: %v", EMPLOYEE_LOG_PREFIX, err))
		return entities.Employee{}, err
	}
	return employee, nil
}

func (s *EmployeeService) RegisterEmployee(request entities.RegisterEmployeeRequest) error {
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error generating salt: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}

	hashedPassword, err := utils.HashPassword(request.Employee.PasswordHash, salt)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error hashing password: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}

	request.Employee.PasswordHash = hashedPassword
	request.Employee.Salt = salt

	err = s.repo.RegisterEmployee(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully registered employee: %v", EMPLOYEE_LOG_PREFIX, request.Employee))
	return nil
}

func (s *EmployeeService) LoginEmployee(ctx context.Context, employeeID, password string) (*entities.Employee, string, error) {
	employee, err := s.repo.LoginEmployee(employeeID, password)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error logging in employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return nil, "", err
	}

	token, err := generateEmployeeJWTToken(employee)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error generating JWT token: %v", EMPLOYEE_LOG_PREFIX, err))
		return nil, "", err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully logged in employee: %s", EMPLOYEE_LOG_PREFIX, employeeID))
	return employee, token, nil
}

func generateEmployeeJWTToken(employee *entities.Employee) (string, error) {
	claims := jwt.MapClaims{
		"employee_id": employee.ID,
		"exp":         time.Now().Add(EMPLOYEE_TOKEN_EXPIRY_DURATION).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecretKey))
}
