package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"ams-service/internal/utils"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

type EmployeeRepositoryImpl struct {
	db *sql.DB
}

func NewEmployeeRepositoryImpl(db *sql.DB) secondary.EmployeeRepository {
	return &EmployeeRepositoryImpl{db: db}
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
		if err == sql.ErrNoRows {
			log.Error().Str("employee_id", request.EmployeeID).Msg("Employee not found")
			return entities.Employee{}, fmt.Errorf("employee not found")
		}
		log.Error().Err(err).Str("employee_id", request.EmployeeID).Msg("Error getting employee by employee_id")
		return entities.Employee{}, err
	}

	return employee, nil
}

func (r *EmployeeRepositoryImpl) RegisterEmployee(request entities.RegisterEmployeeRequest) error {
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
		request.Employee.EmployeeID,
		request.Employee.Name,
		request.Employee.Surname,
		request.Employee.Email,
		request.Employee.Phone,
		request.Employee.Address,
		request.Employee.Gender,
		request.Employee.BirthDate,
		request.Employee.HireDate,
		request.Employee.Position,
		request.Employee.Role,
		request.Employee.Salary,
		request.Employee.Status,
		request.Employee.EmergencyContact,
		request.Employee.EmergencyPhone,
		request.Employee.ProfileImageURL,
		request.Employee.PasswordHash,
		request.Employee.Salt,
	)
	if err != nil {
		log.Error().Err(err).Str("employee_id", request.Employee.EmployeeID).Msg("Error registering employee")
		return err
	}

	log.Info().Str("employee_id", request.Employee.EmployeeID).Msg("Successfully registered employee")
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
		if err == sql.ErrNoRows {
			log.Error().Str("employee_id", employeeID).Msg("Employee not found")
			return nil, fmt.Errorf("employee not found")
		}
		log.Error().Err(err).Str("employee_id", employeeID).Msg("Error logging in employee")
		return nil, err
	}

	isValid, err := utils.VerifyPassword(password, employee.PasswordHash, employee.Salt)
	if err != nil {
		log.Error().Err(err).Str("employee_id", employeeID).Msg("Error verifying password")
		return nil, err
	}
	if !isValid {
		log.Error().Str("employee_id", employeeID).Msg("Invalid password")
		return nil, fmt.Errorf("invalid password")
	}

	log.Info().Str("employee_id", employeeID).Msg("Successfully logged in employee")
	return &employee, nil
}
