package mongodriver

import (
	"time"

	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Entity object
type UserRegister struct {
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

type UserLogin struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password,omitempty" bson:"password"`
}

func responseModel(user UserRegister) domain.UserResponse {
	return domain.UserResponse{
		ID:        user.ID.Hex(),
		UserName:  user.UserName,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
}

func userRegisterEntityToModel(user UserRegister) domain.UserRegister {
	return domain.UserRegister{
		ID:              user.ID.Hex(),
		UserName:        user.UserName,
		Email:           user.Email,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Role:            user.Role,
		CreatedAt:       user.CreatedAt,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
		Verified:        user.Verified,
	}
}
