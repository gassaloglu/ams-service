package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"
	"ams-service/internal/utils"

	"github.com/rs/zerolog/log"
)

type UserService struct {
	repo  secondary.UserRepository
	token primary.TokenService
}

func NewUserService(repo secondary.UserRepository, token primary.TokenService) primary.UserService {
	return &UserService{repo: repo, token: token}
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

	token, err := s.token.CreateUserToken(user)
	if err != nil {
		log.Error().Err(err).Str("username", username).Msg("Error generating user auth token")
		return nil, "", err
	}

	log.Info().Str("username", username).Msg("Successfully logged in user")
	return user, token, nil
}
