package repository

import "github.com/monkeydnoya/hiraishin-auth/pkg/domain"

type AuthRepository interface {
	GetUserById(id string) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	GetUserByUsername(username string) (domain.User, error)
}
