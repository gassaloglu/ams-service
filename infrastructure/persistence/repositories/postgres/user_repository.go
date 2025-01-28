package postgres

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"database/sql"
	"fmt"
)

var USER_LOG_PREFIX string = "user_repository.go"

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) ports.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) RegisterUser(user entities.User) error {
	middlewares.LogInfo(fmt.Sprintf("%s - Registering user: %v", USER_LOG_PREFIX, user))

	query := `
        INSERT INTO users (name, surname, username, email, password_hash, phone, gender, birth_date, role, last_login, created_at, updated_at, last_password_change)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
    `

	_, err := r.db.Exec(query, user.Name, user.Surname, user.Username, user.Email, user.PasswordHash, user.Phone, user.Gender, user.BirthDate, user.Role, user.LastLogin, user.CreatedAt, user.UpdatedAt, user.LastPasswordChange)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering user: %v", USER_LOG_PREFIX, err))
		return err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully registered user: %v", USER_LOG_PREFIX, user))
	return nil
}
