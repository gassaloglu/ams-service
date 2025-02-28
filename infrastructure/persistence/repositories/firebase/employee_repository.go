package firebase

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"context"
	"fmt"
	"time"

	"firebase.google.com/go/v4/db"
)

var EMPLOYEE_LOG_PREFIX string = "employee_repository.go"

type EmployeeRepositoryImpl struct {
	client *db.Client
	ctx    context.Context
}

func NewEmployeeRepositoryImpl(client *db.Client, ctx context.Context) ports.EmployeeRepository {
	return &EmployeeRepositoryImpl{client: client, ctx: ctx}
}

func (r *EmployeeRepositoryImpl) RegisterEmployee(request entities.RegisterEmployeeRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Adding new employee", EMPLOYEE_LOG_PREFIX))

	// Check if an employee with the same national ID already exists
	ref := r.client.NewRef("employees")
	var existingEmployee map[string]interface{}
	err := ref.OrderByChild("employee_id").EqualTo(request.Employee.EmployeeID).Get(r.ctx, &existingEmployee)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error checking for existing employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}
	if len(existingEmployee) > 0 {
		middlewares.LogError(fmt.Sprintf("%s - Employee with national ID %s already exists", EMPLOYEE_LOG_PREFIX, request.Employee.EmployeeID))
		return fmt.Errorf("employee with national ID %s already exists", request.Employee.EmployeeID)
	}

	// Create a new employee
	newEmployee := map[string]interface{}{
		"id":                request.Employee.ID,
		"employee_id":       request.Employee.EmployeeID,
		"name":             request.Employee.Name,
		"surname":          request.Employee.Surname,
		"email":            request.Employee.Email,
		"phone":            request.Employee.Phone,
		"address":          request.Employee.Address,
		"gender":           request.Employee.Gender,
		"birth_date":       request.Employee.BirthDate.Format(time.RFC3339),
		"hire_date":        request.Employee.HireDate.Format(time.RFC3339),
		"position":         request.Employee.Position,
		"department":       request.Employee.Department,
		"salary":           request.Employee.Salary,
		"status":           request.Employee.Status,
		"emergency_contact": request.Employee.EmergencyContact,
		"emergency_phone":   request.Employee.EmergencyPhone,
		"profile_image_url": request.Employee.ProfileImageURL,
		"created_at":       time.Now().Format(time.RFC3339),
		"updated_at":       time.Now().Format(time.RFC3339),
	}

	// Save the new employee to the database
	employeeRef := ref.Child(fmt.Sprintf("%d", request.Employee.ID))
	if err := employeeRef.Set(r.ctx, newEmployee); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error adding employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully added employee with national ID: %s", EMPLOYEE_LOG_PREFIX, request.Employee.EmployeeID))
	return nil
}

func (r *EmployeeRepositoryImpl) GetEmployeeByNationalID(request entities.GetEmployeeByNationalIDRequest) (entities.Employee, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying employee by National ID: %s", EMPLOYEE_LOG_PREFIX, request.EmployeeID))

	ref := r.client.NewRef("employees")
	var employees map[string]entities.Employee
	err := ref.OrderByChild("employee_id").EqualTo(request.EmployeeID).Get(r.ctx, &employees)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying employee by National ID: %v", EMPLOYEE_LOG_PREFIX, err))
		return entities.Employee{}, err
	}

	if len(employees) == 0 {
		middlewares.LogInfo(fmt.Sprintf("%s - Employee with National ID %s not found", EMPLOYEE_LOG_PREFIX, request.EmployeeID))
		return entities.Employee{}, fmt.Errorf("employee with National ID %s not found", request.EmployeeID)
	}

	// Extract the first employee (since employee_id is unique)
	var employee entities.Employee
	for _, emp := range employees {
		employee = emp
		break
	}

	return employee, nil
}

func (r *EmployeeRepositoryImpl) GetEmployeeByID(request entities.GetEmployeeByIDRequest) (entities.Employee, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying employee by ID: %d", EMPLOYEE_LOG_PREFIX, request.ID))

	ref := r.client.NewRef("employees")
	var employee entities.Employee
	err := ref.Child(fmt.Sprintf("%d", request.ID)).Get(r.ctx, &employee)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error querying employee by ID: %v", EMPLOYEE_LOG_PREFIX, err))
		return entities.Employee{}, err
	}

	if employee.ID == 0 {
		middlewares.LogInfo(fmt.Sprintf("%s - Employee with ID %d not found", EMPLOYEE_LOG_PREFIX, request.ID))
		return entities.Employee{}, fmt.Errorf("employee with ID %d not found", request.ID)
	}

	return employee, nil
}
