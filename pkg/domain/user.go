package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	ID        string    `json:"id,omitempty"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Role      []string  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type DBResponse struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName       string             `json:"firstname" bson:"firstname"`
	LastName        string             `json:"lastname" bson:"lastname"`
	UserName        string             `json:"username" bson:"username"`
	Email           string             `json:"email" bson:"email"`
	Password        string             `json:"password" bson:"password"`
	PasswordConfirm string             `json:"passwordConfirm" bson:"passwordConfirm"`
	Role            []string           `json:"role"`
	Verified        bool               `json:"verified" bson:"verified"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
}

type UserRegister struct {
	FirstName       string    `json:"firstname"`
	LastName        string    `json:"lastname" validate:"required"`
	UserName        string    `json:"username" validate:"required"`
	Email           string    `json:"email" validate:"required"`
	Password        string    `json:"password" validate:"required,min=8"`
	PasswordConfirm string    `json:"passwordConfirm" validate:"required"`
	Role            []string  `json:"role"`
	Verified        bool      `json:"verified"`
	CreatedAt       time.Time `json:"created_at"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type Token struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

type ValidateToken struct {
	AccessToken string `json:"accesstoken"`
	UserID      string `json:"userid"`
	IsExpired   bool   `json:"isexpired"`
}

type TokenExpire struct {
	Message string `json:"expire"`
}
