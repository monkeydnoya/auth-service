package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	utils "github.com/monkeydnoya/hiraishin-auth/internal/domain/utils"
	"github.com/monkeydnoya/hiraishin-auth/pkg/config"
	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
)

func (s Server) DeserializeUser() fiber.Handler {
	return s.Service.DeserializeUser()
}

func (s Server) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := s.Service.GetUserById(id)

	if err != nil {
		return c.SendStatus(400)
	}

	return c.Status(200).JSON(user)
}

func (s Server) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := s.Service.GetUserById(email)

	if err != nil {
		return c.SendStatus(400)
	}

	return c.Status(200).JSON(user)
}

func (s Server) GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	user, err := s.Service.GetUserById(username)

	if err != nil {
		return c.SendStatus(400)
	}

	return c.Status(200).JSON(user)
}

func (s Server) SignUp(c *fiber.Ctx) error {
	var user domain.UserRegister

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(400)
	}

	if user.Password != user.PasswordConfirm {
		return c.SendStatus(400)
	}

	userInfo, err := s.Service.RegisterUser(user)
	if err != nil {
		return c.SendStatus(400)
	}

	return c.Status(200).JSON(userInfo)
}

func (s Server) SignIn(c *fiber.Ctx) error {
	var credentials domain.UserLogin

	if err := c.BodyParser(&credentials); err != nil {
		return c.SendStatus(400)
	}

	tokens, err := s.Service.LogIn(credentials)
	if err != nil {
		return c.SendStatus(400)
	}

	// Rethink: Using cookie values instead of returned JSON

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "access_token",
	// 	Value:    access_token,
	// 	Expires:  time.Now().Add(15 * time.Minute),
	// 	HTTPOnly: true,
	// 	SameSite: "lax",
	// })

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "refresh_token",
	// 	Value:    refresh_token,
	// 	Expires:  time.Now().Add(15 * time.Minute),
	// 	HTTPOnly: true,
	// 	SameSite: "lax",
	// })

	return c.Status(200).JSON(tokens)
}

func (s Server) LogoutUser(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Minute),
		HTTPOnly: true,
		SameSite: "lax",
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Minute),
		HTTPOnly: true,
		SameSite: "lax",
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Minute),
		HTTPOnly: true,
		SameSite: "lax",
	})

	return ctx.Status(200).JSON("Status")
}

func (s Server) GetMe(ctx *fiber.Ctx) error {
	currentUser := ctx.Locals("currentUser")
	return ctx.Status(200).JSON(currentUser)
}

func (s Server) ValidateToken(ctx *fiber.Ctx) error {
	// Research:
	// 1. Is correct to do validate process in handler?
	// 2. Return 200 status when token expire

	var accessToken string
	token := ctx.Get("Authorization")
	tokenString := fmt.Sprint(token)
	fields := strings.Fields(tokenString)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	} else if token != "" {
		accessToken = token
	}

	if accessToken == "null" {
		return ctx.SendStatus(401)
	}

	sub, err := utils.ValidateToken(accessToken, config.Config("ACCESS_TOKEN_PUBLIC_KEY"))
	userID := fmt.Sprintf("%s", sub)

	if err != nil {
		if err.Error() == "Token is expired" {
			validation := domain.ValidateToken{
				AccessToken: token,
				UserID:      userID,
				IsExpired:   true,
			}
			return ctx.Status(200).JSON(validation)
		}
	}

	validation := domain.ValidateToken{
		AccessToken: token,
		UserID:      userID,
		IsExpired:   false,
	}

	return ctx.Status(200).JSON(validation)
}

func (s Server) RefreshToken(ctx *fiber.Ctx) error {
	// Research:
	// Some function is not correct

	var refreshToken string
	token := ctx.Get("Authorization")
	tokenString := fmt.Sprint(token)
	fields := strings.Fields(tokenString)

	if len(fields) != 0 && fields[0] == "Bearer" {
		refreshToken = fields[1]
	} else if token != "" {
		refreshToken = token
	}

	if refreshToken == "null" {
		return ctx.SendStatus(401)
	}

	sub, _ := utils.ValidateToken(refreshToken, config.Config("REFRESH_TOKEN_PUBLIC_KEY"))
	userID := fmt.Sprintf("%s", sub)

	accessTokenExpire, _ := time.ParseDuration(config.Config("ACCESS_TOKEN_EXPIRED_IN"))
	accessToken, err := utils.CreateToken(accessTokenExpire, userID, config.Config("ACCESS_TOKEN_PRIVATE_KEY"))
	if err != nil {
		return ctx.SendStatus(400)
	}

	tokens := domain.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return ctx.Status(200).JSON(tokens)
}
