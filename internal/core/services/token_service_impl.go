package services

import (
	"ams-service/internal/core/entities"
	"ams-service/internal/ports/primary"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var UserTokenExpiryDuration = time.Hour * 72
var EmployeeTokenExpiryDuration = time.Hour * 72

type userClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

type employeeClaims struct {
	EmployeeID uint   `json:"employee_id"`
	Role       string `json:"role"`
	jwt.StandardClaims
}

type TokenServiceImpl struct {
	secret  string
	keyFunc jwt.Keyfunc
}

func NewTokenService(secret string) primary.TokenService {
	return &TokenServiceImpl{
		secret: secret,
		keyFunc: func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secret), nil
		},
	}
}

func (service *TokenServiceImpl) CreateUserToken(user *entities.User) (string, error) {
	claims := userClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(UserTokenExpiryDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(service.secret))
}

func (service *TokenServiceImpl) CreateEmployeeToken(employee *entities.Employee) (string, error) {
	claims := employeeClaims{
		employee.ID,
		employee.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(EmployeeTokenExpiryDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(service.secret))
}

func (service *TokenServiceImpl) ValidateToken(tokenString string) error {
	_, err := service.parseStandardClaims(tokenString)
	return err
}

func (service *TokenServiceImpl) ValidateUserToken(tokenString string) error {
	_, err := service.parseUserClaims(tokenString)
	return err
}

func (service *TokenServiceImpl) ValidateEmployeeToken(tokenString string) error {
	_, err := service.parseEmployeeClaims(tokenString)
	return err
}

func (service *TokenServiceImpl) ValidateRole(tokenString string, allowedRoles []string) error {
	claims, err := service.parseEmployeeClaims(tokenString)
	if err != nil {
		return err
	}

	for _, role := range allowedRoles {
		if claims.Role == role {
			return nil
		}
	}

	return errors.New("unauthorized role")
}

func (service *TokenServiceImpl) parseStandardClaims(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, service.keyFunc)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid && claims.Valid() == nil {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (service *TokenServiceImpl) parseUserClaims(tokenString string) (*userClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, service.keyFunc)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*userClaims); ok && token.Valid && claims.Valid() == nil {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (service *TokenServiceImpl) parseEmployeeClaims(tokenString string) (*employeeClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &employeeClaims{}, service.keyFunc)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*employeeClaims); ok && token.Valid && claims.Valid() == nil {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
