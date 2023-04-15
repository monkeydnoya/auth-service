package auth

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/monkeydnoya/hiraishin-auth/internal/domain/utils"
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

func (s Service) LogIn(user domain.UserLogin) (domain.Token, error) {
	response, err := s.DAO.LogIn(user)
	if err != nil {
		configuration.Logger.Errorw("login error:", err)
		return domain.Token{}, err
	}
	return response, nil
}

func (s Service) DeserializeUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var accessToken string
		cookie := ctx.Cookies("access_token")
		authorizationHeader := ctx.Get("Authorization")
		authHeaderString := fmt.Sprint(authorizationHeader)
		fields := strings.Fields(authHeaderString)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if cookie != "" {
			accessToken = cookie
		}
		sub, err := utils.ValidateToken(accessToken, configuration.Config("ACCESS_TOKEN_PUBLIC_KEY"))
		if err != nil {
			configuration.Logger.Infow("validation error: validate token:", err)
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		user, err := s.GetUserById(fmt.Sprint(sub))
		if err != nil {
			configuration.Logger.Infow("validation error: user not found:", err)
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		ctx.Locals("currentUser", user)

		// TODO: Actualize pass cookie
		ctx.Cookie(&fiber.Cookie{
			Name:     "email",
			Value:    user.Email,
			HTTPOnly: true,
			SameSite: "lax",
		})
		ctx.Cookie(&fiber.Cookie{
			Name:     "username",
			Value:    user.UserName,
			HTTPOnly: true,
			SameSite: "lax",
		})

		return ctx.Next()
	}
}
