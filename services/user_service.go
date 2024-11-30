package services

import (
	"errors"
	"messenger-backend/data-access"
	"messenger-backend/models"
	"messenger-backend/utils"
	"os"
	"time"
	"fmt"

	"gopkg.in/gomail.v2"
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

func IsEmailInUse(email string) (bool, error) {
	_, err := dataaccess.FindUserByEmail(email)
	if err != nil {
		if err.Error() == "user not found" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func IsUsernameInUse(username string) (bool, error) {
 _, err := dataaccess.FindUserByUsername(username)
 if err != nil {
	if err.Error() == "user not found" {
		return false, nil
	}
	return false, err
 }
 return false, nil
}

func SendVerificationEmail(email string, code string) error {
	gmailEmail := os.Getenv("GMAIL_EMAIL")
	gmailPassword := os.Getenv("GMAIL_PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")

	m := gomail.NewMessage()
	m.SetHeader("From", gmailEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Verification Code")
	m.SetBody("text/plain", "Your verification code is: "+code)

	d := gomail.NewDialer(smtpServer, 587, gmailEmail, gmailPassword)

	return d.DialAndSend(m)
}

func ProcessVerification(email string) error {
  code := utils.GenerateVerificationCode()

  expiration := time.Now().Add(10 * time.Minute)
  err := dataaccess.StoreVerificationCode(email, code, expiration)
  if err != nil {
    return errors.New("failed to store verification code123")
  }

	fmt.Printf("email: %v, code: %v", email, code)
  err = SendVerificationEmail(email, code)
  if err != nil {
		fmt.Printf("%v", err)
    return errors.New("failed to send verification code")
  }

  return nil
}