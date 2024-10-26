package services

import (
	"messenger-backend/models"
	"messenger-backend/data-access"
	"errors"
)

// CreateUser: Creates a new user and saves it to the database
func CreateUser(user models.User) error {
	return dataaccess.SaveUser(user)
}

// GetUserByEmail: Fetches a user by their email
func GetUserByEmail(email string) (models.User, error) {
	user, err := dataaccess.FindUserByEmail(email)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

// GetUserById: Fetches a user by their ID
func GetUserById(userID uint) (models.User, error) {
	user, err := dataaccess.FindUserByID(userID)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

