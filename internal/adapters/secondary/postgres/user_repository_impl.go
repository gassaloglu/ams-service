package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) secondary.UserRepository {
	db.AutoMigrate(&entities.User{})
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(user *entities.User) (*entities.User, error) {
	clone := *user
	result := r.db.Create(&clone)
	return &clone, result.Error
}

func (r *UserRepositoryImpl) CreateAll(users []entities.User) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&users)
		return result.Error
	})
}

func (r *UserRepositoryImpl) GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *UserRepositoryImpl) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}
