package auth

import (
	"github.com/monkeydnoya/hiraishin-auth/internal/data/repository"
	"github.com/monkeydnoya/hiraishin-auth/internal/service"
)

type Service struct {
	DAO repository.AuthRepository
}

func NewService(auth Service) (service.AuthService, error) {
	service := Service{
		DAO: auth.DAO,
	}
	return service, nil
}
