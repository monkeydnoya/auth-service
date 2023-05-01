package token

import (
	"github.com/golang-jwt/jwt"
	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
)

type Token interface {
	CreateToken(payload interface{}, token *domain.Token) (*domain.Token, error)
	ValidateToken(token domain.Token) (jwt.MapClaims, error)
}
