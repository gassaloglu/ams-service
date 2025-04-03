package postgres

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/secondary"
	"ams-service/internal/utils"
	"database/sql"
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) secondary.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) RegisterUser(user entities.User) error {
	log.Info().
		Str("name", user.Name).
		Str("username", user.Username).
		Str("email", user.Email).
		Msg("Registering user")

	query := `
        INSERT INTO users (
            name, surname, username, email, password_hash, salt, phone, gender, birth_date, 
            last_login, created_at, updated_at, last_password_change
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, DEFAULT, DEFAULT, $11
        )
    `

	_, err := r.db.Exec(
		query,
		user.Name,
		user.Surname,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.Salt,
		user.Phone,
		user.Gender,
		user.BirthDate,
		user.LastLogin,
		user.LastPasswordChange,
	)
	if err != nil {
		log.Error().
			Err(err).
			Str("username", user.Username).
			Msg("Error registering user")
		return fmt.Errorf("failed to register user: %w", err)
	}

	log.Info().
		Str("username", user.Username).
		Msg("Successfully registered user")
	return nil
}

func (r *UserRepositoryImpl) LoginUser(username, password string) (*entities.User, error) {
	log.Info().Str("username", username).Msg("Logging in user")

	query := `
        SELECT id, name, surname, username, email, password_hash, salt, phone, gender, birth_date, 
               last_login, created_at, updated_at, last_password_change 
        FROM users 
        WHERE username = $1
    `
	row := r.db.QueryRow(query, username)

	user, err := r.scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error().Str("username", username).Msg("User not found")
			return nil, fmt.Errorf("user not found")
		}
		log.Error().Err(err).Str("username", username).Msg("Error retrieving user")
		return nil, err
	}

	if err := r.verifyPassword(password, user); err != nil {
		log.Error().Str("username", username).Msg(err.Error())
		return nil, err
	}

	log.Info().Str("username", username).Msg("Successfully logged in user")
	return user, nil
}

// Helper function to scan a user from a database row
func (r *UserRepositoryImpl) scanUser(row *sql.Row) (*entities.User, error) {
	var user entities.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Salt,
		&user.Phone,
		&user.Gender,
		&user.BirthDate,
		&user.LastLogin,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastPasswordChange,
	)
	if err != nil {
		return nil, err
	}
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
