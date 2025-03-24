package services

import (
	"ams-service/config"
	"ams-service/core/entities"
	"ams-service/middlewares"
	"ams-service/utils"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var USER_LOG_PREFIX string = "user_service.go"
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
		middlewares.LogError(fmt.Sprintf("%s - Error generating salt: %v", USER_LOG_PREFIX, err))
		return err
	}

	hashedPassword, err := utils.HashPassword(user.PasswordHash, salt)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error hashing password: %v", USER_LOG_PREFIX, err))
		return err
	}

	user.PasswordHash = hashedPassword
	user.Salt = salt

	err = s.repo.RegisterUser(user)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering user: %v", USER_LOG_PREFIX, err))
		return err
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully registered user: %v", USER_LOG_PREFIX, user))
	return nil
}

func (s *UserService) LoginUser(username, password string) (*entities.User, string, error) {
	user, err := s.repo.LoginUser(username, password)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error logging in user: %v", USER_LOG_PREFIX, err))
		return nil, "", err
	}

	token, err := generateJWTToken(user)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error generating JWT token: %v", USER_LOG_PREFIX, err))
		return nil, "", err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Successfully logged in user: %s", USER_LOG_PREFIX, username))
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
