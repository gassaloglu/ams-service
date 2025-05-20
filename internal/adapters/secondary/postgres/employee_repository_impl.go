package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"

	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRepositoryImpl(db *gorm.DB) secondary.EmployeeRepository {
	db.AutoMigrate(&entities.Employee{})
	return &EmployeeRepositoryImpl{db: db}
}

func (r *EmployeeRepositoryImpl) CreateAll(employee []entities.Employee) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&employee)
		return result.Error
	})
}

func (r *EmployeeRepositoryImpl) FindAll() ([]entities.Employee, error) {
	var employees []entities.Employee
	result := r.db.Find(&employees)
	return employees, result.Error
}

func (r *EmployeeRepositoryImpl) FindByNationalId(nationalId string) (*entities.Employee, error) {
	var employee entities.Employee
	result := r.db.First(&employee, &entities.Employee{NationalID: nationalId})
	return &employee, result.Error
}

func (r *EmployeeRepositoryImpl) Create(employee *entities.Employee) (*entities.Employee, error) {
	var clone = *employee
	result := r.db.Create(&clone)
	return &clone, result.Error
}
