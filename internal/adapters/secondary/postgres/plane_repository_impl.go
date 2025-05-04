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

func (r *PlaneRepositoryImpl) FindAll(req entities.GetAllPlanesRequest) ([]entities.Plane, error) {
	var planes []entities.Plane
	result := r.db.Find(&planes)
	return planes, result.Error
}

func (r *PlaneRepositoryImpl) Create(plane *entities.Plane) (*entities.Plane, error) {
	clone := *plane
	result := r.db.Create(&clone)
	return &clone, result.Error
}
