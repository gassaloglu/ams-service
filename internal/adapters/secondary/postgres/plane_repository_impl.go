package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"

	"gorm.io/gorm"
)

type PlaneRepositoryImpl struct {
	db *gorm.DB
}

func NewPlaneRepositoryImpl(db *gorm.DB) secondary.PlaneRepository {
	db.AutoMigrate(&entities.Plane{})
	return &PlaneRepositoryImpl{db: db}
}

func (r *PlaneRepositoryImpl) GetAllPlanes(req entities.GetAllPlanesRequest) ([]entities.Plane, error) {
	var planes []entities.Plane
	result := r.db.Find(&planes)
	return planes, result.Error
}

func (r *PlaneRepositoryImpl) AddPlane(request entities.AddPlaneRequest) error {
	result := r.db.Create(&request)
	return result.Error
}

func (r *PlaneRepositoryImpl) SetPlaneStatus(request entities.SetPlaneStatusRequest) error {
	result := r.db.Model(&entities.Plane{}).
		Where("registration", request.PlaneRegistration).
		Update("status", request.Status)

	return result.Error
}

func (r *PlaneRepositoryImpl) GetPlaneByRegistration(request entities.GetPlaneByRegistrationRequest) (entities.Plane, error) {
	var plane entities.Plane
	result := r.db.Model(&plane).
		Where("registration", request.Registration).
		Find(&plane)

	return plane, result.Error
}
