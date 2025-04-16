package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"ams-service/internal/utils"
	"database/sql"
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
)

type EmployeeRepositoryImpl struct {
	db *sql.DB
}

func NewEmployeeRepositoryImpl(db *sql.DB) secondary.EmployeeRepository {
	return &EmployeeRepositoryImpl{db: db}
}

// Helper function to scan an employee from a database row
func (r *EmployeeRepositoryImpl) scanEmployee(row *sql.Row) (*entities.Employee, error) {
	var employee entities.Employee
	err := row.Scan(
		&employee.ID,
		&employee.EmployeeID,
		&employee.Name,
		&employee.Surname,
		&employee.Email,
		&employee.Phone,
		&employee.Address,
		&employee.Gender,
		&employee.BirthDate,
		&employee.HireDate,
		&employee.Position,
		&employee.Role,
		&employee.Salary,
		&employee.Status,
		&employee.EmergencyContact,
		&employee.EmergencyPhone,
		&employee.ProfileImageURL,
		&employee.CreatedAt,
		&employee.UpdatedAt,
		&employee.PasswordHash,
		&employee.Salt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("employee not found")
		}
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error) {
	log.Info().Str("employee_id", request.EmployeeID).Msg("Getting employee by employee_id")

	query := `
        SELECT id, employee_id, name, surname, email, phone, address, gender, birth_date, hire_date, 
               position, role, salary, status, emergency_contact, emergency_phone, 
               profile_image_url, created_at, updated_at, password_hash, salt
        FROM employees 
        WHERE employee_id = $1
    `
	row := r.db.QueryRow(query, request.EmployeeID)

	employee, err := r.scanEmployee(row)
	if err != nil {
		log.Error().Err(err).Str("employee_id", request.EmployeeID).Msg("Error getting employee by employee_id")
		return entities.Employee{}, err
	}

	return *employee, nil
}

func (r *EmployeeRepositoryImpl) RegisterEmployee(employee entities.Employee) error {
	log.Info().Msg("Registering new employee")

	query := `
        INSERT INTO employees (
            employee_id, name, surname, email, phone, address, gender, birth_date, hire_date, position, 
            role, salary, status, emergency_contact, emergency_phone, profile_image_url, 
            password_hash, salt, created_at, updated_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, DEFAULT, DEFAULT
        )
    `
	_, err := r.db.Exec(
		query,
		employee.EmployeeID,
		employee.Name,
		employee.Surname,
		employee.Email,
		employee.Phone,
		employee.Address,
		employee.Gender,
		employee.BirthDate,
		employee.HireDate,
		employee.Position,
		employee.Role,
		employee.Salary,
		employee.Status,
		employee.EmergencyContact,
		employee.EmergencyPhone,
		employee.ProfileImageURL,
		employee.PasswordHash,
		employee.Salt,
	)
	if err != nil {
		log.Error().Err(err).Str("employee_id", employee.EmployeeID).Msg("Error registering employee")
		return err
	}

	log.Info().Str("employee_id", employee.EmployeeID).Msg("Successfully registered employee")
	return nil
}

func (r *EmployeeRepositoryImpl) LoginEmployee(employeeID, password string) (*entities.Employee, error) {
	log.Info().Str("employee_id", employeeID).Msg("Logging in employee")

	query := `
        SELECT id, employee_id, name, surname, email, phone, address, gender, birth_date, hire_date, 
               position, role, salary, status, emergency_contact, emergency_phone, 
               profile_image_url, created_at, updated_at, password_hash, salt
        FROM employees 
        WHERE employee_id = $1
    `
	row := r.db.QueryRow(query, employeeID)

	employee, err := r.scanEmployee(row)
	if err != nil {
		log.Error().Err(err).Str("employee_id", employeeID).Msg("Error logging in employee")
		return nil, err
	}

	if err := r.verifyPassword(password, employee); err != nil {
		log.Error().Str("employee_id", employeeID).Msg(err.Error())
		return nil, err
	}

	log.Info().Str("employee_id", employeeID).Msg("Successfully logged in employee")
	return employee, nil
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
