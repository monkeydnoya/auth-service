package domain

const (
	Access  = "access"
	Refresh = "refresh"
)

type Token struct {
	Token     string
	TokenType string
}

type JWTTokenResponse struct {
	AccessToken  string
	RefreshToken string
}
