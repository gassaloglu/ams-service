package services

import (
	"ams-service/internal/config"
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"
	"ams-service/internal/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
)

var EmployeeTokenExpiryDuration = time.Hour * 72

type EmployeeServiceImpl struct {
	repo secondary.EmployeeRepository
}

func NewEmployeeService(repo secondary.EmployeeRepository) primary.EmployeeService {
	return &EmployeeServiceImpl{repo: repo}
}

func (s *EmployeeServiceImpl) GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error) {
	employee, err := s.repo.GetEmployeeByID(request)
	if err != nil {
		log.Error().Err(err).Str("employee_id", request.EmployeeID).Msg("Error getting employee by employee_id")
		return entities.Employee{}, err
	}
	return employee, nil
}

func (s *EmployeeServiceImpl) RegisterEmployee(request entities.RegisterEmployeeRequest) error {
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		log.Error().Err(err).Msg("Error generating salt")
		return err
	}

	hashedPassword, err := utils.HashPassword(request.Employee.PasswordHash, salt)
	if err != nil {
		log.Error().Err(err).Msg("Error hashing password")
		return err
	}

	request.Employee.PasswordHash = hashedPassword
	request.Employee.Salt = salt

	err = s.repo.RegisterEmployee(request)
	if err != nil {
		log.Error().Err(err).Msg("Error registering employee")
		return err
	}
	log.Info().Interface("employee", request.Employee).Msg("Successfully registered employee")
	return nil
}

func (s *EmployeeServiceImpl) LoginEmployee(employeeID, password string) (*entities.Employee, string, error) {
	employee, err := s.repo.LoginEmployee(employeeID, password)
	if err != nil {
		log.Error().Err(err).Str("employee_id", employeeID).Msg("Error logging in employee")
		return nil, "", err
	}

	token, err := generateEmployeeJWTToken(employee)
	if err != nil {
		log.Error().Err(err).Str("employee_id", employeeID).Msg("Error generating JWT token")
		return nil, "", err
	}

	log.Info().Str("employee_id", employeeID).Msg("Successfully logged in employee")
	return employee, token, nil
}

func generateEmployeeJWTToken(employee *entities.Employee) (string, error) {
	claims := jwt.MapClaims{
		"employee_id": employee.ID,
		"exp":         time.Now().Add(EmployeeTokenExpiryDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecretKey))
}
