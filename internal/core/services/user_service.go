package services

import (
	"ams-service/internal/config"
	"ams-service/internal/core/entities"
	"ams-service/internal/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
)

var TOKEN_EXPIRY_DURATION time.Duration = time.Hour * 72

type UserRepository interface {
	RegisterUser(user entities.User) error
	LoginUser(username, password string) (*entities.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user entities.User) error {
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		log.Error().Err(err).Msg("Error generating salt")
		return err
	}

	hashedPassword, err := utils.HashPassword(user.PasswordHash, salt)
	if err != nil {
		log.Error().Err(err).Msg("Error hashing password")
		return err
	}

	user.PasswordHash = hashedPassword
	user.Salt = salt

	err = s.repo.RegisterUser(user)
	if err != nil {
		log.Error().Err(err).Msg("Error registering user")
		return err
	}
	log.Info().Str("username", user.Username).Msg("Successfully registered user")
	return nil
}

func (s *UserService) LoginUser(username, password string) (*entities.User, string, error) {
	user, err := s.repo.LoginUser(username, password)
	if err != nil {
		log.Error().Err(err).Str("username", username).Msg("Error logging in user")
		return nil, "", err
	}

	token, err := generateJWTToken(user)
	if err != nil {
		log.Error().Err(err).Str("username", username).Msg("Error generating JWT token")
		return nil, "", err
	}

	log.Info().Str("username", username).Msg("Successfully logged in user")
	return user, token, nil
}

func generateJWTToken(user *entities.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(TOKEN_EXPIRY_DURATION).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecretKey))
}
