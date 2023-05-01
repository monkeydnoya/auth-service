package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
)

func (s Server) GetUserById(c *fiber.Ctx) error {
	// TODO: Change parsing configration from params to Path
	id := c.Params("id")
	user, err := s.Service.GetUserById(id)

	if err != nil {
		return c.SendStatus(400)
	}

	return c.Status(200).JSON(user)
}

func (s Server) GetUserByEmail(c *fiber.Ctx) error {
	// TODO: Change parsing configration from params to Path
	email := c.Params("email")
	user, err := s.Service.GetUserByEmail(email)

	if err != nil {
		return c.SendStatus(400)
	}

	return c.Status(200).JSON(user)
}

func (s Server) GetUserByUsername(c *fiber.Ctx) error {
	// TODO: Change parsing configration from params to Path
	username := c.Params("username")
	user, err := s.Service.GetUserByUsername(username)

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
		return c.SendStatus(401)
	}

	tokens, err := s.Service.LogIn(credentials)
	if err != nil {
		return c.SendStatus(401)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    tokens.AccessToken,
		HTTPOnly: true,
		SameSite: "lax",
		Secure:   true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		HTTPOnly: true,
		SameSite: "lax",
		Secure:   true,
	})

	return c.SendStatus(200)
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

	return ctx.SendStatus(200)
}

func (s Server) ValidateToken(ctx *fiber.Ctx) error {
	var accessToken string = ctx.Cookies("access_token")
	token := domain.Token{Token: accessToken, TokenType: domain.Access}

	user, err := s.Service.ValidateToken(token)
	if err != nil {
		return ctx.SendStatus(403)
	}

	return ctx.Status(200).JSON(user)
}

func (s Server) RefreshToken(ctx *fiber.Ctx) error {
	var refreshToken string = ctx.Cookies("refresh_token")
	var err error

	tokens := &domain.JWTTokenResponse{
		RefreshToken: refreshToken,
	}
	tokens, err = s.Service.RefreshToken(tokens)
	if err != nil {
		return ctx.SendStatus(401)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    tokens.AccessToken,
		HTTPOnly: true,
		SameSite: "lax",
		Secure:   true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		HTTPOnly: true,
		SameSite: "lax",
		Secure:   true,
	})

	return ctx.SendStatus(200)
}
