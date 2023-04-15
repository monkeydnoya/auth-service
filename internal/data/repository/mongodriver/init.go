package mongodriver

import (
	"github.com/monkeydnoya/hiraishin-auth/internal/data/repository"
	configuration "github.com/monkeydnoya/hiraishin-auth/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Client   *mongo.Client
	Database *mongo.Database
}

type AuthDAO struct {
	DB *mongo.Database
}

func (c Config) Init() (repository.AuthRepository, error) {
	configuration.Logger.Infow("Repository for mongo driver initialized")
	auth := AuthDAO{
		DB: c.Database,
	}

	return auth, nil
}
