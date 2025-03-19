package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"ams-service/middlewares"

	"golang.org/x/crypto/pbkdf2"
)

const PASSWORD_LOG_PREFIX string = "password.go"

// GenerateSalt generates a new salt for password hashing
func GenerateSalt(size int) (string, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Failed to generate salt: %v", PASSWORD_LOG_PREFIX, err))
		return "", fmt.Errorf("failed to generate salt: %v", err)
	}
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully generated salt", PASSWORD_LOG_PREFIX))
	return base64.StdEncoding.EncodeToString(salt), nil
}

// HashPassword hashes a password using PBKDF2 with the given salt
func HashPassword(password, salt string) (string, error) {
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Failed to decode salt: %v", PASSWORD_LOG_PREFIX, err))
		return "", fmt.Errorf("failed to decode salt: %v", err)
	}

	hash := pbkdf2.Key([]byte(password), saltBytes, 10000, 32, sha256.New)
	hashedPassword := base64.StdEncoding.EncodeToString(hash)
	middlewares.LogInfo(fmt.Sprintf("%s - Successfully hashed password", PASSWORD_LOG_PREFIX))
	return hashedPassword, nil
}

// VerifyPassword verifies a password against a given hash and salt
func VerifyPassword(password, hash, salt string) (bool, error) {
	hashedPassword, err := HashPassword(password, salt)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error hashing password: %v", PASSWORD_LOG_PREFIX, err))
		return false, err
	}
	isValid := hashedPassword == hash
	if isValid {
		middlewares.LogInfo(fmt.Sprintf("%s - Password verification successful", PASSWORD_LOG_PREFIX))
	} else {
		middlewares.LogError(fmt.Sprintf("%s - Password verification failed", PASSWORD_LOG_PREFIX))
	}
	return isValid, nil
}
