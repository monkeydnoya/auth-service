package main

import (
	"time"

	"github.com/monkeydnoya/hiraishin-auth/internal/controller"
	"github.com/monkeydnoya/hiraishin-auth/internal/controller/fiber/handler"
	"github.com/monkeydnoya/hiraishin-auth/internal/data/config/mongodb"
	"github.com/monkeydnoya/hiraishin-auth/internal/data/repository/mongodriver"
	"github.com/monkeydnoya/hiraishin-auth/internal/domain/token/jwttoken"
	"github.com/monkeydnoya/hiraishin-auth/internal/service/auth"
	configuration "github.com/monkeydnoya/hiraishin-auth/pkg/config"
)

func main() {
	config := &mongodb.Config{
		Host:     configuration.Config("MONGO_URI"),
		Port:     configuration.Config("PORT"),
		DBName:   configuration.Config("DB"),
		Username: configuration.Config("DB_USERNAME"),
		Password: configuration.Config("DB_PASSWORD"),
	}

	mongoconnect, err := config.Connect()
	if err != nil {
		// Rethink:
		// Handle in any error in config.Connect() returns os.Exit()
		// Handle error for the future
		configuration.Logger.Infow("MongoDB connection error", err)
	}

	authDAO, err := mongodriver.Config{Client: mongoconnect.Client, Database: mongoconnect.DbConnection}.Init()
	if err != nil {
		// Repository interface do not return error
		// Handle error for the future
		configuration.Logger.Infow("MongoDB driver error", err)
	}

	expiration, err := time.ParseDuration(configuration.Config("ACCESS_TOKEN_EXPIRED_IN"))
	if err != nil {
		panic(err)
	}
	accessToken := jwttoken.Token{
		PrivateKey: configuration.Config("ACCESS_TOKEN_PRIVATE_KEY"),
		PublicKey:  configuration.Config("ACCESS_TOKEN_PUBLIC_KEY"),
		Expiration: expiration,
	}

	expiration, err = time.ParseDuration(configuration.Config("REFRESH_TOKEN_EXPIRED_IN"))
	if err != nil {
		panic(err)
	}
	refreshToken := jwttoken.Token{
		PrivateKey: configuration.Config("REFRESH_TOKEN_PRIVATE_KEY"),
		PublicKey:  configuration.Config("REFRESH_TOKEN_PUBLIC_KEY"),
		Expiration: expiration,
	}

	jwtToken, err := jwttoken.NewJWTToken(accessToken, refreshToken)
	if err != nil {
		configuration.Logger.Infow("Token initialize error", err)
	}

	service, err := auth.NewService(auth.Service{
		DAO:      authDAO,
		JWTToken: jwtToken,
	})
	if err != nil {
		// Repository interface do not return error
		// Handle error for the future
		configuration.Logger.Infow("Service initialize error", err)
	}

	server := handler.NewServer(service)
	controller.StartServer(server)
}
