package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"ams-service/internal/ports/secondary"
	"ams-service/internal/utils"
	"errors"
)

type UserService struct {
	repo  secondary.UserRepository
	token primary.TokenService
}

func NewUserService(repo secondary.UserRepository, token primary.TokenService) primary.UserService {
	return &UserService{repo: repo, token: token}
}

func (s *UserService) Register(request *entities.RegisterUserRequest) (string, error) {
	user, err := mapRegisterUserRequestToUserEntity(request)
	if err != nil {
		return "", err
	}

	user, err = s.repo.CreateUser(user)
	if err != nil {
		return "", err
	}

	token, err := s.token.CreateUserToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	isValid, err := utils.VerifyPassword(password, user.PasswordHash, user.Salt)
	if err != nil || !isValid {
		return "", errors.New("could not verify password")
	}

	token, err := s.token.CreateUserToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func mapRegisterUserRequestToUserEntity(request *entities.RegisterUserRequest) (*entities.User, error) {
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(request.Password, salt)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Name:         request.Name,
		Surname:      request.Surname,
		Email:        request.Email,
		Phone:        request.Phone,
		Gender:       request.Gender,
		BirthDate:    request.BirthDate,
		PasswordHash: hashedPassword,
		Salt:         salt,
	}

	return user, nil
}
