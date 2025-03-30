package postgres

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/utils"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

var EMPLOYEE_LOG_PREFIX string = "employee_repository.go"

type EmployeeRepositoryImpl struct {
	db *sql.DB
}

func NewEmployeeRepositoryImpl(db *sql.DB) ports.EmployeeRepository {
	return &EmployeeRepositoryImpl{db: db}
}

func (r *EmployeeRepositoryImpl) GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error) {
	log.Info().Uint("id", request.ID).Msg("Getting employee by ID")

	query := `SELECT id, employee_id, name, surname, email, phone, address, gender, birth_date, hire_date, position, department, salary, status, manager_id, emergency_contact, emergency_phone, profile_image_url, created_at, updated_at FROM employees WHERE id = $1`
	row := r.db.QueryRow(query, request.ID)

	var employee entities.Employee
	err := row.Scan(&employee.ID, &employee.EmployeeID, &employee.Name, &employee.Surname, &employee.Email, &employee.Phone, &employee.Address, &employee.Gender, &employee.BirthDate, &employee.HireDate, &employee.Position, &employee.Department, &employee.Salary, &employee.Status, &employee.ManagerID, &employee.EmergencyContact, &employee.EmergencyPhone, &employee.ProfileImageURL, &employee.CreatedAt, &employee.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Uint("id", request.ID).Msg("Error getting employee by ID")
		return entities.Employee{}, err
	}

	return employee, nil
}

func (r *EmployeeRepositoryImpl) RegisterEmployee(request entities.RegisterEmployeeRequest) error {
	log.Info().Msg("Registering new employee")

	// Check if an employee with the same ID already exists
	query := `SELECT id FROM employees WHERE employee_id = $1`
	row := r.db.QueryRow(query, request.Employee.ID)
	var existingID uint
	err := row.Scan(&existingID)
	if err == nil {
		log.Error().Uint("employee_id", request.Employee.ID).Msg("Employee with ID already exists")
		return fmt.Errorf("employee with ID %d already exists", request.Employee.ID)
	} else if err != sql.ErrNoRows {
		log.Error().Err(err).Msg("Error checking for existing employee")
		return err
	}

	// Log the gender value before insertion
	log.Info().Interface("employee", request.Employee).Msg("Employee details")

	query = `INSERT INTO employees (employee_id, name, surname, email, phone, address, gender, birth_date, hire_date, position, department, salary, status, manager_id, emergency_contact, emergency_phone, profile_image_url, password_hash, salt, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)`
	_, err = r.db.Exec(query, request.Employee.EmployeeID, request.Employee.Name, request.Employee.Surname, request.Employee.Email, request.Employee.Phone, request.Employee.Address, request.Employee.Gender, request.Employee.BirthDate, request.Employee.HireDate, request.Employee.Position, request.Employee.Department, request.Employee.Salary, request.Employee.Status, request.Employee.ManagerID, request.Employee.EmergencyContact, request.Employee.EmergencyPhone, request.Employee.ProfileImageURL, request.Employee.PasswordHash, request.Employee.Salt, request.Employee.CreatedAt, request.Employee.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Str("name", request.Employee.Name).Str("surname", request.Employee.Surname).Msg("Error registering employee")
		return err
	}
	return nil
}

func (r *EmployeeRepositoryImpl) LoginEmployee(employeeID, password string) (*entities.Employee, error) {
	log.Info().Str("employee_id", employeeID).Msg("Logging in employee")

	query := `SELECT id, employee_id, name, surname, email, password_hash, salt FROM employees WHERE employee_id = $1`
	row := r.db.QueryRow(query, employeeID)

	var employee entities.Employee
	err := row.Scan(&employee.ID, &employee.EmployeeID, &employee.Name, &employee.Surname, &employee.Email, &employee.PasswordHash, &employee.Salt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Str("employee_id", employeeID).Msg("Employee not found")
			return nil, fmt.Errorf("employee not found")
		}
		log.Error().Err(err).Str("employee_id", employeeID).Msg("Error logging in employee")
		return nil, err
	}

	isValid, err := utils.VerifyPassword(password, employee.PasswordHash, employee.Salt)
	if err != nil {
		log.Error().Err(err).Str("employee_id", employeeID).Msg("Error verifying password for employee")
		return nil, err
	}
	if !isValid {
		log.Error().Str("employee_id", employeeID).Msg("Invalid password for employee")
		return nil, fmt.Errorf("invalid password")
	}

	log.Info().Str("employee_id", employeeID).Msg("Successfully logged in employee")
	return &employee, nil
}
