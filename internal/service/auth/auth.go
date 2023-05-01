package auth

import (
	pwd "github.com/monkeydnoya/hiraishin-auth/internal/domain/utils"
	configuration "github.com/monkeydnoya/hiraishin-auth/pkg/config"
	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
)

func (s Service) GetUserById(id string) (domain.UserResponse, error) {
	user, err := s.DAO.GetUserById(id)
	if err != nil {
		configuration.Logger.Error("get user: user by id:", err)
		return domain.UserResponse{}, err
	}
	return user, nil
}

func (s Service) GetUserByEmail(email string) (domain.UserResponse, error) {
	user, err := s.DAO.GetUserByEmail(email)
	if err != nil {
		configuration.Logger.Error("get user: user by email:", err)
		return domain.UserResponse{}, err
	}
	return user, nil
}

func (s Service) GetUserByUsername(username string) (domain.UserResponse, error) {
	user, err := s.DAO.GetUserByUsername(username)
	if err != nil {
		configuration.Logger.Error("get user: user by usererrname:", err)
		return domain.UserResponse{}, err
	}
	return user, nil
}

func (s Service) RegisterUser(user domain.UserRegister) (domain.UserResponse, error) {
	response, err := s.DAO.RegisterUser(user)
	if err != nil {
		configuration.Logger.Error("register user:", err)
		return domain.UserResponse{}, err
	}
	return response, nil
}

func (s Service) LogIn(credentials domain.UserLogin) (domain.JWTTokenResponse, error) {
	user, err := s.DAO.GetUser(credentials.Username)
	if err != nil {
		return domain.JWTTokenResponse{}, err
	}

	if err = pwd.VerifyPassword(user.Password, credentials.Password); err != nil {
		return domain.JWTTokenResponse{}, err
	}

	accessToken := &domain.Token{
		TokenType: domain.Access,
	}
	accessToken, err = s.JWTToken.CreateToken(user.ID, accessToken)
	if err != nil {
		return domain.JWTTokenResponse{}, err
	}

	refreshToken := &domain.Token{
		TokenType: domain.Refresh,
	}
	refreshToken, err = s.JWTToken.CreateToken(user.ID, refreshToken)
	if err != nil {
		return domain.JWTTokenResponse{}, err
	}

	jwttoken := domain.JWTTokenResponse{
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
	}

	return jwttoken, nil
}
