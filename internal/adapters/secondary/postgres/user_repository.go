package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"ams-service/internal/utils"
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) secondary.UserRepository {
	db.AutoMigrate(&entities.User{})
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) RegisterUser(user entities.User) error {
	result := r.db.Create(&user)
	return result.Error
}

func (r *UserRepositoryImpl) LoginUser(username, password string) (*entities.User, error) {
	var user entities.User
	result := r.db.Find(&user, &entities.User{Username: username})
	if result.Error != nil {
		return nil, result.Error
	}

	if err := r.verifyPassword(password, &user); err != nil {
		log.Error().Str("username", username).Msg(err.Error())
		return nil, err
	}

	log.Info().Str("username", username).Msg("Successfully logged in user")
	return &user, nil
}

// Helper function to verify a user's password
func (r *UserRepositoryImpl) verifyPassword(password string, user *entities.User) error {
	isValid, err := utils.VerifyPassword(password, user.PasswordHash, user.Salt)
	if err != nil {
		return fmt.Errorf("error verifying password: %w", err)
	}
	if !isValid {
		return fmt.Errorf("invalid password")
	}
	return nil
}
