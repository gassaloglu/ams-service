package postgres

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
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
