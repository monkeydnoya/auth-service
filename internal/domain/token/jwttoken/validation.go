package jwttoken

import (
	"encoding/base64"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
)

func (j jwtToken) ValidateToken(token domain.Token) (jwt.MapClaims, error) {
	var decodedPublicKey []byte
	var err error

	switch token.TokenType {
	case domain.Access:
		decodedPublicKey, err = base64.StdEncoding.DecodeString(j.AccessToken.PublicKey)
		if err != nil {
			return nil, fmt.Errorf("could not decode public key: %w", err)
		}
	case domain.Refresh:
		decodedPublicKey, err = base64.StdEncoding.DecodeString(j.RefreshToken.PublicKey)
		if err != nil {
			return nil, fmt.Errorf("could not decode public key: %w", err)
		}
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return jwt.MapClaims{}, fmt.Errorf("validate: parse key: %w", err)
	}

	parsedToken, err := jwt.Parse(token.Token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}
	return claims, nil
}
