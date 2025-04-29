package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"ams-service/internal/utils"
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRepositoryImpl(db *gorm.DB) secondary.EmployeeRepository {
	db.AutoMigrate(&entities.Employee{})
	return &EmployeeRepositoryImpl{db: db}
}

func (r *EmployeeRepositoryImpl) GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error) {
	var employee entities.Employee
	result := r.db.Find(&employee, &entities.Employee{EmployeeID: request.EmployeeID})
	return employee, result.Error
}

func (r *EmployeeRepositoryImpl) RegisterEmployee(employee entities.Employee) error {
	result := r.db.Create(&employee)
	return result.Error
}

func (r *EmployeeRepositoryImpl) LoginEmployee(employeeID, password string) (*entities.Employee, error) {
	var employee entities.Employee
	result := r.db.Find(&employee, &entities.Employee{EmployeeID: employeeID})
	if result.Error != nil {
		return nil, result.Error
	}

	if err := r.verifyPassword(password, &employee); err != nil {
		log.Error().Str("employee_id", employeeID).Msg(err.Error())
		return nil, err
	}

	log.Info().Str("employee_id", employeeID).Msg("Successfully logged in employee")
	return &employee, nil
}

// Helper function to verify an employee's password
func (r *EmployeeRepositoryImpl) verifyPassword(password string, employee *entities.Employee) error {
	isValid, err := utils.VerifyPassword(password, employee.PasswordHash, employee.Salt)
	if err != nil {
		return fmt.Errorf("error verifying password: %w", err)
	}
	if !isValid {
		return fmt.Errorf("invalid password")
	}
	return nil
}
