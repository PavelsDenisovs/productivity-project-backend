package services

import (
	"messenger-backend/models"
	"messenger-backend/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	Register(user *models.User) (string, string, error) // Returns access and refresh tokens
	Login(email, password string) (string, string, error) // Returns access and refresh tokens
	Logout(refreshToken string) error
	RefreshToken(refreshToken string) (string, string, error) // Returs new access and refresh tokens
	VerifyEmail(email, code string) error
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




func

func (s *userService) CreateUser(user *models.User) error {
	
}