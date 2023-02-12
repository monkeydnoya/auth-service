package mongodriver

import "github.com/monkeydnoya/hiraishin-auth/pkg/domain"

// Entity object
type User struct {
	Id        int64    `json:"id" bson:"id"`
	Username  string   `json:"username" bson:"username"`
	Email     string   `json:"email" bson:"email"`
	FirstName string   `json:"firstname" bson:"firstname"`
	LastName  string   `json:"lastname" bson:"lastname"`
	Role      []string `json:"role" bson:"role"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

func toModel(user User) domain.User {
	return domain.User{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}
}
