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

func (r *EmployeeRepositoryImpl) GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Getting employee by ID: %d", EMPLOYEE_LOG_PREFIX, request.ID))

	ref := r.client.NewRef("employees").Child(fmt.Sprintf("%d", request.ID))
	var employee entities.Employee
	err := ref.Get(r.ctx, &employee)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting employee by ID: %v", EMPLOYEE_LOG_PREFIX, err))
		return entities.Employee{}, err
	}
	return employee, nil
}

func (r *EmployeeRepositoryImpl) RegisterEmployee(request entities.RegisterEmployeeRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Registering new employee", EMPLOYEE_LOG_PREFIX))

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
		"name":              request.Employee.Name,
		"surname":           request.Employee.Surname,
		"email":             request.Employee.Email,
		"phone":             request.Employee.Phone,
		"address":           request.Employee.Address,
		"gender":            request.Employee.Gender,
		"birth_date":        request.Employee.BirthDate.Format(time.RFC3339),
		"hire_date":         request.Employee.HireDate.Format(time.RFC3339),
		"position":          request.Employee.Position,
		"department":        request.Employee.Department,
		"salary":            request.Employee.Salary,
		"status":            request.Employee.Status,
		"emergency_contact": request.Employee.EmergencyContact,
		"emergency_phone":   request.Employee.EmergencyPhone,
		"profile_image_url": request.Employee.ProfileImageURL,
		"created_at":        time.Now().Format(time.RFC3339),
		"updated_at":        time.Now().Format(time.RFC3339),
	}

	employeeRef := ref.Child(fmt.Sprintf("%d", request.Employee.ID))
	if err := employeeRef.Set(r.ctx, newEmployee); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully registered employee with national ID: %s", EMPLOYEE_LOG_PREFIX, request.Employee.EmployeeID))
	return nil
}
