package datagen

import "ams-service/internal/core/services"

type DatagenController struct {
	bankService     services.BankService
	employeeService services.EmployeeServiceImpl
	flightService   services.FlightService
}
