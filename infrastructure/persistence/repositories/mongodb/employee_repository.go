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

func (r *EmployeeRepositoryImpl) RegisterEmployee(request entities.RegisterEmployeeRequest) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Adding new employee", EMPLOYEE_LOG_PREFIX))

	ctx := context.Background()

	// Check if a plane with the same registration already exists
	filter := bson.M{"national_id": request.Employee.ID}
	var existingEmployee entities.Employee
	err := r.collection.FindOne(ctx, filter).Decode(&existingEmployee)
	if err == nil {
		middlewares.LogError(fmt.Sprintf("%s - Employee with national id %s already exists", EMPLOYEE_LOG_PREFIX, request.Employee.EmployeeID))
		return fmt.Errorf("Employee with national id %s already exists", request.Employee.EmployeeID)
	} else if err != mongo.ErrNoDocuments {
		middlewares.LogError(fmt.Sprintf("%s - Error checking for existing employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}

	// Create a new plane document
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
		//ManagerID:        request.Employee.ManagerID,
		//CreatedAt:        request.Employee.CreatedAt,
		//UpdatedAt:        request.Employee.UpdatedAt,
	}

	// Insert the new plane into the collection
	_, err = r.collection.InsertOne(ctx, newEmployee)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error adding employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully added employee with national id: %s", EMPLOYEE_LOG_PREFIX, request.Employee.EmployeeID))
	return nil
}

func (r *EmployeeRepositoryImpl) GetEmployeeByNationalID(request entities.GetEmployeeByNationalIDRequest) (entities.Employee, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying employee by National ID: %s", PASSENGER_LOG_PREFIX, request.EmployeeID))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"national_id": request.EmployeeID}

	var employee entities.Employee
	err := r.collection.FindOne(ctx, filter).Decode(&employee)
	if err != nil {
		middlewares.LogInfo(fmt.Sprintf("%s - Error querying passenger by National ID: %v", EMPLOYEE_LOG_PREFIX, err))
		return entities.Employee{}, err
	}

	return employee, nil
}

func (r *EmployeeRepositoryImpl) GetEmployeeByID(request entities.GetEmployeeByIDRequest) (entities.Employee, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Querying employee by ID: %d", PASSENGER_LOG_PREFIX, request.ID))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": request.ID}

	var employee entities.Employee
	err := r.collection.FindOne(ctx, filter).Decode(&employee)
	if err != nil {
		middlewares.LogInfo(fmt.Sprintf("%s - Error querying passenger by ID: %v", EMPLOYEE_LOG_PREFIX, err))
		return entities.Employee{}, err
	}

	return employee, nil
}
