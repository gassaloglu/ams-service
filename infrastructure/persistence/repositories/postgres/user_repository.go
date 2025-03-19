package postgres

import (
	"ams-service/application/ports"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"ams-service/utils"
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
        INSERT INTO users (name, surname, username, email, password_hash, salt, phone, gender, birth_date, role, last_login, created_at, updated_at, last_password_change)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
    `

	_, err := r.db.Exec(query, user.Name, user.Surname, user.Username, user.Email, user.PasswordHash, user.Salt, user.Phone, user.Gender, user.BirthDate, user.Role, user.LastLogin, user.CreatedAt, user.UpdatedAt, user.LastPasswordChange)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering user: %v", USER_LOG_PREFIX, err))
		return err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully registered user: %v", USER_LOG_PREFIX, user))
	return nil
}

func (r *UserRepositoryImpl) LoginUser(username, password string) (*entities.User, error) {
	middlewares.LogInfo(fmt.Sprintf("%s - Logging in user: %s", USER_LOG_PREFIX, username))

	query := `SELECT id, name, surname, username, email, password_hash, salt FROM users WHERE username = $1`
	row := r.db.QueryRow(query, username)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.Username, &user.Email, &user.PasswordHash, &user.Salt)
	if err != nil {
		if err == sql.ErrNoRows {
			middlewares.LogError(fmt.Sprintf("%s - User not found: %s", USER_LOG_PREFIX, username))
			return nil, fmt.Errorf("user not found")
		}
		middlewares.LogError(fmt.Sprintf("%s - Error logging in user: %v", USER_LOG_PREFIX, err))
		return nil, err
	}

	isValid, err := utils.VerifyPassword(password, user.PasswordHash, user.Salt)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error verifying password for user: %s, error: %v", USER_LOG_PREFIX, username, err))
		return nil, err
	}
	if !isValid {
		middlewares.LogError(fmt.Sprintf("%s - Invalid password for user: %s", USER_LOG_PREFIX, username))
		return nil, fmt.Errorf("invalid password")
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully logged in user: %s", USER_LOG_PREFIX, username))
	return &user, nil
}
