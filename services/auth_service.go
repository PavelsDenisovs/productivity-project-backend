package services

type AuthService interface {
	Login(email, password string) (accessToken, refreshToken string, err error)
	Logout(refreshToken string) error

	RefreshAccessToken(refreshToken string) (newAccessToken, newRefreshToken string, err error)
}