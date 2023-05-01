package repository

import (
	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
)

type AuthRepository interface {
	GetUserById(id string) (domain.UserResponse, error)
	GetUserByEmail(email string) (domain.UserResponse, error)
	GetUserByUsername(username string) (domain.UserResponse, error)

	RegisterUser(user domain.UserRegister) (domain.UserResponse, error)

	GetUser(login string) (domain.UserRegister, error)
}
