package services

import (
	"errors"
	"fmt"
	"productivity-project-backend/models"
	"productivity-project-backend/repository"
	"productivity-project-backend/utils"
)

type AuthService interface {
	Register(email, password string) (*models.User, error)
	Login(email, password string) (*models.User, error)
	VerifyEmail(email, code string) error
	GenerateAndStoreVerificationCode(email string) error
	GetUserByEmail(email string) (*models.User, error)
}

type authService struct {
	userRepo repository.UserRepository
	verificationRepo repository.VerificationRepository
}

func NewAuthService(userRepo repository.UserRepository, verificationRepo repository.VerificationRepository) AuthService {
	return &authService{
		userRepo:         userRepo,
		verificationRepo: verificationRepo,
	}
}

func (s *authService) Register(email, password string) (*models.User, error) {
	existingUser, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		// TODO make errors centralized
		if err.Error() != "user not found" {
			return nil, fmt.Errorf("error checking user existence: %w", err)
		}
	} else if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	if err := utils.ValidatePassword(password); err != nil {
		return nil, fmt.Errorf("failed validation: %w", err)
	}

	if err := utils.ValidateEmail(email); err != nil {
		return nil, fmt.Errorf("failed validation: %w", err)
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        email,
		PasswordHash: hashedPassword,
		IsVerified:   false,
	}
	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return user, nil
}

func (s *authService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsVerified {
		return nil, errors.New("email not verified")
	}

	return user, nil
}

func (s *authService) VerifyEmail(email, code string) error {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	valid, err := s.verificationRepo.VerifyCode(user.ID, code)
	if err != nil {
		return fmt.Errorf("failed to verify code: %w", err)
	}

	if !valid {
		return errors.New("invalid code")
	}

	if err := s.userRepo.MarkEmailAsVerified(email); err != nil {
		return fmt.Errorf("failed to mark email as verified: %w", err)
	}

	return nil
}

func (s *authService) GenerateAndStoreVerificationCode(email string) error {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	code := utils.GenerateVerificationCode()
	
	if err := s.verificationRepo.StoreVerificationCode(user.ID, code); err != nil {
    return fmt.Errorf("failed to store verification code: %w", err)
	}

	if err := utils.SendVerificationEmail(email, code); err != nil {
		return fmt.Errorf("failed to store verification code: %w", err)
	}

	return nil
}

func (s *authService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}