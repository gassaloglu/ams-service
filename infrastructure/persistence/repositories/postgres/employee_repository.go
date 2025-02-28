package postgres

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"database/sql"
	"fmt"
	"time"
)

var EMPLOYEE_LOG_PREFIX string = "employee_repository.go"

type EmployeeRepositoryImpl struct {
	db *sql.DB
}

func NewEmployeeRepositoryImpl(db *sql.DB) ports.EmployeeRepository {
	return &EmployeeRepositoryImpl{db: db}
}

func (r *EmployeeRepositoryImpl) RegisterEmployee(request entities.RegisterEmployeeRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Adding new employee", EMPLOYEE_LOG_PREFIX))

	// Check if an employee with the same national ID already exists
	var existingEmployeeID string
	query := `SELECT employee_id FROM employees WHERE employee_id = $1`
	err := r.db.QueryRow(query, request.Employee.EmployeeID).Scan(&existingEmployeeID)
	if err == nil {
		middlewares.LogError(fmt.Sprintf("%s - Employee with national ID %s already exists", EMPLOYEE_LOG_PREFIX, request.Employee.EmployeeID))
		return fmt.Errorf("employee with national ID %s already exists", request.Employee.EmployeeID)
	} else if err != sql.ErrNoRows {
		middlewares.LogError(fmt.Sprintf("%s - Error checking for existing employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}

	// Insert the new employee into the database
	insertQuery := `
		INSERT INTO employees (
			id, employee_id, name, surname, email, phone, address, gender, birth_date, hire_date,
			position, department, salary, status, emergency_contact, emergency_phone, profile_image_url, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
	`
	_, err = r.db.Exec(
		insertQuery,
		request.Employee.ID,
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
		request.Employee.Department,
		request.Employee.Salary,
		request.Employee.Status,
		request.Employee.EmergencyContact,
		request.Employee.EmergencyPhone,
		request.Employee.ProfileImageURL,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error adding employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully added employee with national ID: %s", EMPLOYEE_LOG_PREFIX, request.Employee.EmployeeID))
	return nil
}

func (r *EmployeeRepositoryImpl) GetEmployeeByNationalID(request entities.GetEmployeeByNationalIDRequest) (entities.Employee, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying employee by National ID: %s", EMPLOYEE_LOG_PREFIX, request.EmployeeID))

	query := `
		SELECT id, employee_id, name, surname, email, phone, address, gender, birth_date, hire_date,
		       position, department, salary, status, emergency_contact, emergency_phone, profile_image_url, created_at, updated_at
		FROM employees
		WHERE employee_id = $1
	`
	var employee entities.Employee
	err := r.db.QueryRow(query, request.EmployeeID).Scan(
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
		&employee.Department,
		&employee.Salary,
		&employee.Status,
		&employee.EmergencyContact,
		&employee.EmergencyPhone,
		&employee.ProfileImageURL,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			middlewares.LogInfo(fmt.Sprintf("%s - Employee with National ID %s not found", EMPLOYEE_LOG_PREFIX, request.EmployeeID))
			return entities.Employee{}, fmt.Errorf("employee with National ID %s not found", request.EmployeeID)
		}
		middlewares.LogError(fmt.Sprintf("%s - Error querying employee by National ID: %v", EMPLOYEE_LOG_PREFIX, err))
		return entities.Employee{}, err
	}

	return employee, nil
}

func (r *EmployeeRepositoryImpl) GetEmployeeByID(request entities.GetEmployeeByIDRequest) (entities.Employee, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying employee by ID: %d", EMPLOYEE_LOG_PREFIX, request.ID))

	query := `
		SELECT id, employee_id, name, surname, email, phone, address, gender, birth_date, hire_date,
		       position, department, salary, status, emergency_contact, emergency_phone, profile_image_url, created_at, updated_at
		FROM employees
		WHERE id = $1
	`
	var employee entities.Employee
	err := r.db.QueryRow(query, request.ID).Scan(
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
		&employee.Department,
		&employee.Salary,
		&employee.Status,
		&employee.EmergencyContact,
		&employee.EmergencyPhone,
		&employee.ProfileImageURL,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			middlewares.LogInfo(fmt.Sprintf("%s - Employee with ID %d not found", EMPLOYEE_LOG_PREFIX, request.ID))
			return entities.Employee{}, fmt.Errorf("employee with ID %d not found", request.ID)
		}
		middlewares.LogError(fmt.Sprintf("%s - Error querying employee by ID: %v", EMPLOYEE_LOG_PREFIX, err))
		return entities.Employee{}, err
	}

	return employee, nil
}
