package jwttoken

import (
	"time"

	"github.com/monkeydnoya/hiraishin-auth/internal/domain/token"
)

type Token struct {
	PrivateKey string
	PublicKey  string
	Expiration time.Duration
}

type jwtToken struct {
	AccessToken  Token
	RefreshToken Token
}

func NewJWTToken(accessToken Token, refreshToken Token) (token.Token, error) {
	jwt := jwtToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return jwt, nil
}
