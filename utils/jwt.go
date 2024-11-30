package utils

import (
	"time"
	"os"
	"github.com/dgrijalva/jwt-go"
	"messenger-backend/models"
)


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