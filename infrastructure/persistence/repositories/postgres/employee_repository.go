package postgres

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"ams-service/utils"
	"database/sql"
	"fmt"
)

var EMPLOYEE_LOG_PREFIX string = "employee_repository.go"

type EmployeeRepositoryImpl struct {
	db *sql.DB
}

func NewEmployeeRepositoryImpl(db *sql.DB) ports.EmployeeRepository {
	return &EmployeeRepositoryImpl{db: db}
}

func (r *EmployeeRepositoryImpl) GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Getting employee by ID: %d", EMPLOYEE_LOG_PREFIX, request.ID))

	query := `SELECT id, employee_id, name, surname, email, phone, address, gender, birth_date, hire_date, position, department, salary, status, manager_id, emergency_contact, emergency_phone, profile_image_url, created_at, updated_at FROM employees WHERE id = $1`
	row := r.db.QueryRow(query, request.ID)

	var employee entities.Employee
	err := row.Scan(&employee.ID, &employee.EmployeeID, &employee.Name, &employee.Surname, &employee.Email, &employee.Phone, &employee.Address, &employee.Gender, &employee.BirthDate, &employee.HireDate, &employee.Position, &employee.Department, &employee.Salary, &employee.Status, &employee.ManagerID, &employee.EmergencyContact, &employee.EmergencyPhone, &employee.ProfileImageURL, &employee.CreatedAt, &employee.UpdatedAt)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting employee by ID: %v", EMPLOYEE_LOG_PREFIX, err))
		return entities.Employee{}, err
	}

	return employee, nil
}

func (r *EmployeeRepositoryImpl) RegisterEmployee(request entities.RegisterEmployeeRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Registering new employee", EMPLOYEE_LOG_PREFIX))

	// Check if an employee with the same ID already exists
	query := `SELECT id FROM employees WHERE employee_id = $1`
	row := r.db.QueryRow(query, request.Employee.ID)
	var existingID uint
	err := row.Scan(&existingID)
	if err == nil {
		middlewares.LogError(fmt.Sprintf("%s - Employee with ID %d already exists", EMPLOYEE_LOG_PREFIX, request.Employee.ID))
		return fmt.Errorf("employee with ID %d already exists", request.Employee.ID)
	} else if err != sql.ErrNoRows {
		middlewares.LogError(fmt.Sprintf("%s - Error checking for existing employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}

	// Log the gender value before insertion
	middlewares.LogInfo(fmt.Sprintf("%s - Employee details: %+v", EMPLOYEE_LOG_PREFIX, request.Employee))

	query = `INSERT INTO employees (employee_id, name, surname, email, phone, address, gender, birth_date, hire_date, position, department, salary, status, manager_id, emergency_contact, emergency_phone, profile_image_url, password_hash, salt, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)`
	_, err = r.db.Exec(query, request.Employee.EmployeeID, request.Employee.Name, request.Employee.Surname, request.Employee.Email, request.Employee.Phone, request.Employee.Address, request.Employee.Gender, request.Employee.BirthDate, request.Employee.HireDate, request.Employee.Position, request.Employee.Department, request.Employee.Salary, request.Employee.Status, request.Employee.ManagerID, request.Employee.EmergencyContact, request.Employee.EmergencyPhone, request.Employee.ProfileImageURL, request.Employee.PasswordHash, request.Employee.Salt, request.Employee.CreatedAt, request.Employee.UpdatedAt)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}
	return nil
}
func (r *EmployeeRepositoryImpl) LoginEmployee(employeeID, password string) (*entities.Employee, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Logging in employee: %s", EMPLOYEE_LOG_PREFIX, employeeID))

	query := `SELECT id, employee_id, name, surname, email, password_hash, salt FROM employees WHERE employee_id = $1`
	row := r.db.QueryRow(query, employeeID)

	var employee entities.Employee
	err := row.Scan(&employee.ID, &employee.EmployeeID, &employee.Name, &employee.Surname, &employee.Email, &employee.PasswordHash, &employee.Salt)
	if err != nil {
		if err == sql.ErrNoRows {
			middlewares.LogError(fmt.Sprintf("%s - Employee not found: %s", EMPLOYEE_LOG_PREFIX, employeeID))
			return nil, fmt.Errorf("employee not found")
		}
		middlewares.LogError(fmt.Sprintf("%s - Error logging in employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return nil, err
	}

	isValid, err := utils.VerifyPassword(password, employee.PasswordHash, employee.Salt)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error verifying password for employee: %s, error: %v", EMPLOYEE_LOG_PREFIX, employeeID, err))
		return nil, err
	}
	if !isValid {
		middlewares.LogError(fmt.Sprintf("%s - Invalid password for employee: %s", EMPLOYEE_LOG_PREFIX, employeeID))
		return nil, fmt.Errorf("invalid password")
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully logged in employee: %s", EMPLOYEE_LOG_PREFIX, employeeID))
	return &employee, nil
}
