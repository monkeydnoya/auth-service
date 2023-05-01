package mongodriver

import (
	"context"
	"errors"
	"strings"
	"time"

	pwd "github.com/monkeydnoya/hiraishin-auth/internal/domain/utils"
	"github.com/monkeydnoya/hiraishin-auth/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (a AuthDAO) GetUserById(id string) (domain.UserResponse, error) {
	user := UserRegister{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.UserResponse{}, err
	}

	err = a.DB.Collection("User").FindOne(ctx, bson.M{"_id": idObjectID}).Decode(&user)
	if err != nil {
		return domain.UserResponse{}, err
	}

	return responseModel(user), nil
}

func (a AuthDAO) GetUserByEmail(email string) (domain.UserResponse, error) {
	user := domain.UserResponse{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := a.DB.Collection("User").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return domain.UserResponse{}, err
	}

	return user, nil
}

func (a AuthDAO) GetUserByUsername(username string) (domain.UserResponse, error) {
	user := UserRegister{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := a.DB.Collection("Users").FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return responseModel(UserRegister{}), err
	}

	return responseModel(user), nil
}

func (a AuthDAO) RegisterUser(user domain.UserRegister) (domain.UserResponse, error) {
	user.CreatedAt = time.Now()
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = true

	hashedPassword, err := pwd.HashPassword(user.Password)
	if err != nil {
		return domain.UserResponse{}, err
	}
	user.Password = hashedPassword

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := a.DB.Collection("User").InsertOne(ctx, &user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return domain.UserResponse{}, errors.New("user with that email already exist")
		}
		return responseModel(UserRegister{}), err
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := a.DB.Collection("User").Indexes().CreateOne(ctx, index); err != nil {
		return domain.UserResponse{}, errors.New("could not create index for email")
	}

	var newUser UserRegister
	err = a.DB.Collection("User").FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&newUser)
	if err != nil {
		return responseModel(UserRegister{}), err
	}

	return responseModel(newUser), nil
}

func (a AuthDAO) GetUser(login string) (domain.UserRegister, error) {
	var err error
	user := UserRegister{}
	loginSplited := strings.Split(login, "@")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if len(loginSplited) > 1 {
		err = a.DB.Collection("User").FindOne(ctx, bson.M{"email": login}).Decode(&user)
	} else {
		err = a.DB.Collection("User").FindOne(ctx, bson.M{"username": login}).Decode(&user)
	}

	if err != nil {
		return domain.UserRegister{}, err
	}

	return userRegisterEntityToModel(user), nil
}
