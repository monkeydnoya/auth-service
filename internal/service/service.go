package service

import (
	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
)

type AuthService interface {
	GetUserById(id string) (domain.UserResponse, error)
	GetUserByEmail(email string) (domain.UserResponse, error)
	GetUserByUsername(username string) (domain.UserResponse, error)

	RegisterUser(user domain.UserRegister) (domain.UserResponse, error)
	LogIn(user domain.UserLogin) (domain.JWTTokenResponse, error)

	ValidateToken(token domain.Token) (domain.UserResponse, error)
	RefreshToken(*domain.JWTTokenResponse) (*domain.JWTTokenResponse, error)
}
