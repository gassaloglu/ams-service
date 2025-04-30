package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSalt(t *testing.T) {
	salt, err := GenerateSalt(16)
	assert.NoError(t, err, "Error should not occur while generating salt")
	assert.NotEmpty(t, salt, "Generated salt should not be empty")
	assert.Equal(t, 24, len(salt), "Encoded salt length should match expected base64 length")
}

func TestHashPassword(t *testing.T) {
	password := "securepassword"
	salt, err := GenerateSalt(16)
	assert.NoError(t, err, "Error should not occur while generating salt")

	hashedPassword, err := HashPassword(password, salt)
	assert.NoError(t, err, "Error should not occur while hashing password")
	assert.NotEmpty(t, hashedPassword, "Hashed password should not be empty")
}

func TestVerifyPassword(t *testing.T) {
	password := "securepassword"
	salt, err := GenerateSalt(16)
	assert.NoError(t, err, "Error should not occur while generating salt")

	hashedPassword, err := HashPassword(password, salt)
	assert.NoError(t, err, "Error should not occur while hashing password")

	// Test valid password
	isValid, err := VerifyPassword(password, hashedPassword, salt)
	assert.NoError(t, err, "Error should not occur while verifying password")
	assert.True(t, isValid, "Password verification should succeed for valid password")

	// Test invalid password
	isValid, err = VerifyPassword("wrongpassword", hashedPassword, salt)
	assert.NoError(t, err, "Error should not occur while verifying password")
	assert.False(t, isValid, "Password verification should fail for invalid password")
}
