package auth

import (
	"github.com/monkeydnoya/hiraishin-auth/internal/data/repository"
	"github.com/monkeydnoya/hiraishin-auth/internal/domain/token"
	"github.com/monkeydnoya/hiraishin-auth/internal/service"
)

type Service struct {
	DAO      repository.AuthRepository
	JWTToken token.Token
}

func NewService(auth Service) (service.AuthService, error) {
	service := Service{
		DAO:      auth.DAO,
		JWTToken: auth.JWTToken,
	}
	return service, nil
}
