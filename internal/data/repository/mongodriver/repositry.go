package mongodriver

import (
	"context"
	"fmt"
	"time"

	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
)

// Interface realization
func (a AuthDAO) GetUserByEmail(email string) (domain.User, error) {
	user := User{}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := a.DB.Collection("Users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		fmt.Println("Error when finding in database")
		return domain.User{}, err
	}

	return toModel(user), nil
}

func (a AuthDAO) GetUserById(id string) (domain.User, error) {
	user := User{}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := a.DB.Collection("Users").FindOne(ctx, bson.M{"id": id}).Decode(&user)
	if err != nil {
		fmt.Println("Error when finding in database")
		return toModel(User{}), err
	}

	return toModel(user), nil
}

func (a AuthDAO) GetUserByUsername(username string) (domain.User, error) {
	user := User{}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := a.DB.Collection("Users").FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		fmt.Println("Error when finding in database")
		return domain.User{}, err
	}

	return toModel(user), nil
}
