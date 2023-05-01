package jwttoken

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
)

func (j jwtToken) CreateToken(payload interface{}, token *domain.Token) (*domain.Token, error) {
	var decodedPrivateKey []byte
	var tokenExpire time.Duration
	var err error

	switch token.TokenType {
	case domain.Access:
		decodedPrivateKey, err = base64.StdEncoding.DecodeString(j.AccessToken.PrivateKey)
		if err != nil {
			return token, fmt.Errorf("could not decode public key: %w", err)
		}
		tokenExpire = j.AccessToken.Expiration
	case domain.Refresh:
		decodedPrivateKey, err = base64.StdEncoding.DecodeString(j.RefreshToken.PrivateKey)
		if err != nil {
			return token, fmt.Errorf("could not decode public key: %w", err)
		}
		tokenExpire = j.RefreshToken.Expiration
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return token, fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(tokenExpire).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	newToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return token, fmt.Errorf("create: sign token: %w", err)
	}

	token.Token = newToken

	return token, nil
}
