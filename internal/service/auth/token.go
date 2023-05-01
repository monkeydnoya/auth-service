package auth

import (
	"fmt"

	configuration "github.com/monkeydnoya/hiraishin-auth/pkg/config"
	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
)

func (s Service) ValidateToken(token domain.Token) (domain.UserResponse, error) {
	claims, err := s.JWTToken.ValidateToken(token)
	if err != nil {
		configuration.Logger.Error("jwt token: token validation error:", err)
		return domain.UserResponse{}, err
	}
	userID := fmt.Sprintf("%s", claims["sub"])
	user, err := s.DAO.GetUserById(userID)
	if err != nil {
		configuration.Logger.Error("jwt token: user not found:", err)
		return domain.UserResponse{}, err
	}
	return user, nil
}

func (s Service) RefreshToken(tokens *domain.JWTTokenResponse) (*domain.JWTTokenResponse, error) {
	refreshToken := domain.Token{
		Token:     tokens.RefreshToken,
		TokenType: domain.Refresh,
	}
	claims, err := s.JWTToken.ValidateToken(refreshToken)
	if err != nil {
		configuration.Logger.Error("jwt token: token validation error:", err)
		return &domain.JWTTokenResponse{}, err
	}
	userID := fmt.Sprintf("%s", claims["sub"])
	user, err := s.DAO.GetUserById(userID)
	if err != nil {
		configuration.Logger.Error("jwt token: user not found:", err)
		return &domain.JWTTokenResponse{}, err
	}

	accessToken := &domain.Token{
		TokenType: domain.Access,
	}
	accessToken, err = s.JWTToken.CreateToken(user.ID, accessToken)
	if err != nil {
		configuration.Logger.Error("jwt token: new token create error:", err)
		return &domain.JWTTokenResponse{}, err
	}

	tokens.AccessToken = accessToken.Token
	return tokens, nil
}
