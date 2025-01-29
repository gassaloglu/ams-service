package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

// GenerateSalt generates a new salt for password hashing
func GenerateSalt(size int) (string, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %v", err)
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// HashPassword hashes a password using PBKDF2 with the given salt
func HashPassword(password, salt string) (string, error) {
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return "", fmt.Errorf("failed to decode salt: %v", err)
	}

	hash := pbkdf2.Key([]byte(password), saltBytes, 10000, 32, sha256.New)
	return base64.StdEncoding.EncodeToString(hash), nil
}

// VerifyPassword verifies a password against a given hash and salt
func VerifyPassword(password, salt, hash string) (bool, error) {
	hashedPassword, err := HashPassword(password, salt)
	if err != nil {
		return false, err
	}
	return hashedPassword == hash, nil
}
