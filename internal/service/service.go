package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
)

type AuthService interface {
	GetUserById(id string) (domain.UserResponse, error)
	GetUserByEmail(email string) (domain.UserResponse, error)
	GetUserByUsername(username string) (domain.UserResponse, error)

	RegisterUser(user domain.UserRegister) (domain.UserResponse, error)
	LogIn(user domain.UserLogin) (domain.Token, error)

	DeserializeUser() fiber.Handler
}
