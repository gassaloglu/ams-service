package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/pbkdf2"
)

// GenerateSalt generates a new salt for password hashing
func GenerateSalt(size int) (string, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate salt")
		return "", fmt.Errorf("failed to generate salt: %v", err)
	}

	return base64.StdEncoding.EncodeToString(salt), nil
}

// HashPassword hashes a password using PBKDF2 with the given salt
func HashPassword(password, salt string) (string, error) {
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		log.Error().Err(err).Msg("Failed to decode salt")
		return "", fmt.Errorf("failed to decode salt: %v", err)
	}

	bytes := []byte(password)
	hash := pbkdf2.Key(bytes, saltBytes, 10000, 32, sha256.New)
	hashedPassword := base64.StdEncoding.EncodeToString(hash)
	return hashedPassword, nil
}

// VerifyPassword verifies a password against a given hash and salt
func VerifyPassword(password, hash, salt string) (bool, error) {
	hashedPassword, err := HashPassword(password, salt)
	if err != nil {
		log.Error().Err(err).Msg("Error hashing password")
		return false, err
	}
	isValid := hashedPassword == hash
	if isValid {
		log.Info().Msg("Password verification successful")
	} else {
		log.Error().Msg("Password verification failed")
	}
	return isValid, nil
}
