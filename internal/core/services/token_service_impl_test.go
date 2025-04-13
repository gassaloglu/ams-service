package services_test

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/core/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTokenService(t *testing.T) {
	service := services.NewTokenService("secret")
	assert.NotNil(t, service)
}

func TestTokenServiceImpl_CreateUserToken(t *testing.T) {
	assert := assert.New(t)
	service := services.NewTokenService("secret")
	user := &entities.User{ID: 1}

	token, err := service.CreateUserToken(user)

	assert.NoError(err)
	assert.NotEmpty(token)
}

func TestTokenServiceImpl_CreateEmployeeToken(t *testing.T) {
	assert := assert.New(t)
	service := services.NewTokenService("secret")
	employee := &entities.Employee{ID: 1, Role: "admin"}

	token, err := service.CreateEmployeeToken(employee)

	assert.NoError(err)
	assert.NotEmpty(token)
}

func TestTokenServiceImpl_ValidateToken(t *testing.T) {
	assert := assert.New(t)
	service := services.NewTokenService("secret")

	user := &entities.User{ID: 1}
	employee := &entities.Employee{ID: 1, Role: "admin"}

	userToken, _ := service.CreateUserToken(user)
	employeeToken, _ := service.CreateEmployeeToken(employee)

	userTokenErr := service.ValidateToken(userToken)
	employeeTokenErr := service.ValidateToken(employeeToken)

	assert.NoError(userTokenErr)
	assert.NoError(employeeTokenErr)
}

func TestTokenServiceImpl_ValidateUserToken(t *testing.T) {
	assert := assert.New(t)
	service := services.NewTokenService("secret")
	user := &entities.User{ID: 1}

	userToken, _ := service.CreateUserToken(user)
	userTokenErr := service.ValidateToken(userToken)

	assert.NoError(userTokenErr)
}

func TestTokenServiceImpl_ValidateEmployeeToken(t *testing.T) {
	assert := assert.New(t)
	service := services.NewTokenService("secret")
	employee := &entities.Employee{ID: 1, Role: "admin"}

	employeeToken, _ := service.CreateEmployeeToken(employee)
	employeeTokenErr := service.ValidateToken(employeeToken)

	assert.NoError(employeeTokenErr)
}

func TestTokenServiceImpl_ValidateRole(t *testing.T) {
	assert := assert.New(t)
	service := services.NewTokenService("secret")

	tests := []struct {
		name         string
		employee     *entities.Employee
		allowedRoles []string
		wantErr      bool
	}{
		{
			name:         "Valid role",
			employee:     &entities.Employee{ID: 1, Role: "admin"},
			allowedRoles: []string{"admin", "user"},
			wantErr:      false,
		},
		{
			name:         "Valid ground_services role",
			employee:     &entities.Employee{ID: 5, Role: "ground_services"},
			allowedRoles: []string{"ground_services", "passenger_services"},
			wantErr:      false,
		},
		{
			name:         "Multiple allowed roles including matching role",
			employee:     &entities.Employee{ID: 6, Role: "flight_planner"},
			allowedRoles: []string{"admin", "hr", "flight_planner", "passenger_services", "ground_services"},
			wantErr:      false,
		},
		{
			name:         "Invalid role - admin not allowed",
			employee:     &entities.Employee{ID: 7, Role: "admin"},
			allowedRoles: []string{"hr", "flight_planner"},
			wantErr:      true,
		},
		{
			name:         "Empty allowed roles list",
			employee:     &entities.Employee{ID: 12, Role: "admin"},
			allowedRoles: []string{},
			wantErr:      true,
		},
		{
			name:         "Unknown role not in system",
			employee:     &entities.Employee{ID: 14, Role: "unknown_role"},
			allowedRoles: []string{"admin", "hr", "flight_planner", "passenger_services", "ground_services"},
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			employeeToken, _ := service.CreateEmployeeToken(tt.employee)
			employeeTokenErr := service.ValidateRole(employeeToken, tt.allowedRoles)

			if tt.wantErr {
				assert.Error(employeeTokenErr)
			} else {
				assert.NoError(employeeTokenErr)
			}
		})
	}
}
