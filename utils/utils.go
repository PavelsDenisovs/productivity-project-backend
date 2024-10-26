package utils

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	"os"
	"github.com/dgrijalva/jwt-go"
	"messenger-backend/models"
)

// HashPassword: Hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash: Compares a bcrypt hashed password with its possible plaintext equivalent
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// JWT secret key 
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT: Generates JWT token for the user
func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}