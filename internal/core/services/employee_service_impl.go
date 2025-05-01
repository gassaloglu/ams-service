package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"
	"ams-service/internal/utils"
	"errors"
)

type EmployeeServiceImpl struct {
	repo  secondary.EmployeeRepository
	token primary.TokenService
}

func NewEmployeeService(repo secondary.EmployeeRepository, token primary.TokenService) primary.EmployeeService {
	return &EmployeeServiceImpl{repo: repo, token: token}
}

func (s *EmployeeServiceImpl) FindAll() ([]entities.Employee, error) {
	return s.repo.FindAll()
}

func (s *EmployeeServiceImpl) RegisterAll(requests []entities.RegisterEmployeeRequest) error {
	var employees []entities.Employee

	for _, request := range requests {
		employee, err := mapRegisterEmployeeRequestToEmployeeEntity(&request)

		if err != nil {
			return err
		}

		employees = append(employees, *employee)
	}

	return s.repo.CreateAll(employees)
}

func (s *EmployeeServiceImpl) Register(request *entities.RegisterEmployeeRequest) (string, error) {
	employee, err := mapRegisterEmployeeRequestToEmployeeEntity(request)
	if err != nil {
		return "", err
	}

	employee, err = s.repo.Create(employee)
	if err != nil {
		return "", err
	}

	token, err := s.token.CreateEmployeeToken(employee)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *EmployeeServiceImpl) Login(request *entities.LoginEmployeeRequest) (string, error) {
	employee, err := s.repo.FindByNationalId(request.NationalID)
	if err != nil {
		return "", err
	}

	isValid, err := utils.VerifyPassword(request.Password, employee.PasswordHash, employee.Salt)
	if err != nil || !isValid {
		return "", errors.New("could not verify password")
	}

	token, err := s.token.CreateEmployeeToken(employee)
	if err != nil {
		return "", err
	}

	return token, nil
}

func mapRegisterEmployeeRequestToEmployeeEntity(request *entities.RegisterEmployeeRequest) (*entities.Employee, error) {
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(request.Password, salt)
	if err != nil {
		return nil, err
	}

	employee := &entities.Employee{
		NationalID:   request.NationalID,
		Name:         request.Name,
		Surname:      request.Surname,
		Email:        request.Email,
		Phone:        request.Phone,
		Gender:       request.Gender,
		BirthDate:    request.BirthDate,
		PasswordHash: hashedPassword,
		Salt:         salt,
		Title:        request.Title,
		Role:         request.Role,
	}

	return employee, nil
}
