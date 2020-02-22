package common

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSecretPassword(password string) string {
	bytes := sha256.Sum256([]byte(password))
	return hex.EncodeToString(bytes[:])
}

func ValidatePassword(password string, secretPassword string) bool {
	return secretPassword == GetSecretPassword(password)
}