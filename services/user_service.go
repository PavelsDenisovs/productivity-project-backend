package services

import (
	"errors"
	"fmt"
	"messenger-backend/models"
	"messenger-backend/repository"
	"messenger-backend/utils"
)

type UserService interface {
	Register(email, password string, userData models.User) error
	ResendVerificationCode(email string) error
	VerifyEmail(email, code string) error

	GetUserProfile(userID string) (*models.User, error)
	UpdateUserProfile(userID string, updateData models.UserUpdate)

	SearchUsers(query string) ([]models.User, error)
}

type userService struct {
	userRepo         repository.UserRepository
	jwtRepo			 repository.JWTTokenRepository
	verificationRepo repository.VerificationRepository
}

func NewUserService(userRepo repository.UserRepository, jwtRepo repository.JWTTokenRepository, verificationRepo repository.VerificationRepository) UserService {
	return &userService{
		userRepo: userRepo,
		jwtRepo: jwtRepo,
		verificationRepo: verificationRepo,
	}
}

func (s *userService) Register(email, password string, userData models.User) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err) 
	}
	userData.PasswordHash = hashedPassword
	userData.Email = email

	err = s.userRepo.CreateUser(&userData)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) ResendVerificationCode(email string) error {
	code := utils.GenerateVerificationCode()
	err := utils.SendVerificationEmail(email, code)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) VerifyEmail(email, code string) error {
	isVerified, err := s.verificationRepo.VerifyCode(email, code)
	if err != nil {
		return err
	}
	if !isVerified {
		return errors.New("invalid or expired verification code")
	}
	return nil
}