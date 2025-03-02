package mongodb

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var EMPLOYEE_LOG_PREFIX string = "employee_repository.go"

type EmployeeRepositoryImpl struct {
	collection *mongo.Collection
}

func NewEmployeeRepositoryImpl(client *mongo.Client, dbName, collectionName string) ports.EmployeeRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &EmployeeRepositoryImpl{collection: collection}
}

func (r *EmployeeRepositoryImpl) GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Getting employee by ID: %d", EMPLOYEE_LOG_PREFIX, request.ID))

	filter := bson.M{"id": request.ID}
	var employee entities.Employee
	err := r.collection.FindOne(context.Background(), filter).Decode(&employee)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting employee by ID: %v", EMPLOYEE_LOG_PREFIX, err))
		return entities.Employee{}, err
	}
	return employee, nil
}

func (r *EmployeeRepositoryImpl) RegisterEmployee(request entities.RegisterEmployeeRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Registering new employee", EMPLOYEE_LOG_PREFIX))

	ctx := context.Background()

	// Check if an employee with the same national ID already exists
	filter := bson.M{"employee_id": request.Employee.EmployeeID}
	var existingEmployee entities.Employee
	err := r.collection.FindOne(ctx, filter).Decode(&existingEmployee)
	if err == nil {
		middlewares.LogError(fmt.Sprintf("%s - Employee with national ID %s already exists", EMPLOYEE_LOG_PREFIX, request.Employee.EmployeeID))
		return fmt.Errorf("employee with national ID %s already exists", request.Employee.EmployeeID)
	} else if err != mongo.ErrNoDocuments {
		middlewares.LogError(fmt.Sprintf("%s - Error checking for existing employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}

	// Create a new employee
	newEmployee := entities.Employee{
		ID:               request.Employee.ID,
		EmployeeID:       request.Employee.EmployeeID,
		Name:             request.Employee.Name,
		Surname:          request.Employee.Surname,
		Email:            request.Employee.Email,
		Phone:            request.Employee.Phone,
		Address:          request.Employee.Address,
		Gender:           request.Employee.Gender,
		BirthDate:        request.Employee.BirthDate,
		HireDate:         request.Employee.HireDate,
		Position:         request.Employee.Position,
		Department:       request.Employee.Department,
		Salary:           request.Employee.Salary,
		Status:           request.Employee.Status,
		EmergencyContact: request.Employee.EmergencyContact,
		EmergencyPhone:   request.Employee.EmergencyPhone,
		ProfileImageURL:  request.Employee.ProfileImageURL,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	_, err = r.collection.InsertOne(ctx, newEmployee)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}
	return nil
}
