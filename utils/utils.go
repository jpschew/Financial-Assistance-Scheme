package utils

import (
	fjwt "FinancialAssistanceScheme/middleware/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims = fjwt.UserClaims

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword checks if a given password matches the hashed password
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
