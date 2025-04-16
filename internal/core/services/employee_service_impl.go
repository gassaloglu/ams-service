package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"
	"ams-service/internal/utils"

	"github.com/rs/zerolog/log"
)

type EmployeeServiceImpl struct {
	repo  secondary.EmployeeRepository
	token primary.TokenService
}

func NewEmployeeService(repo secondary.EmployeeRepository, token primary.TokenService) primary.EmployeeService {
	return &EmployeeServiceImpl{repo: repo, token: token}
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
	for _, employee := range request {
		salt, err := utils.GenerateSalt(16)
		if err != nil {
			log.Error().Err(err).Msg("Error generating salt")
			return err
		}

		hashedPassword, err := utils.HashPassword(employee.PasswordHash, salt)
		if err != nil {
			log.Error().Err(err).Msg("Error hashing password")
			return err
		}

		employee.PasswordHash = hashedPassword
		employee.Salt = salt

		err = s.repo.RegisterEmployee(employee)
		if err != nil {
			log.Error().Err(err).Msg("Error registering employee")
			return err
		}
		log.Info().Interface("employee", employee).Msg("Successfully registered employee")
	}
	return nil
}

func (s *EmployeeServiceImpl) LoginEmployee(employeeID, password string) (*entities.Employee, string, error) {
	employee, err := s.repo.LoginEmployee(employeeID, password)
	if err != nil {
		log.Error().Err(err).Str("employee_id", employeeID).Msg("Error logging in employee")
		return nil, "", err
	}

	token, err := s.token.CreateEmployeeToken(employee)
	if err != nil {
		log.Error().Err(err).Str("employee_id", employeeID).Msg("Error generating employee auth token")
		return nil, "", err
	}

	log.Info().Str("employee_id", employeeID).Msg("Successfully logged in employee")
	return employee, token, nil
}
