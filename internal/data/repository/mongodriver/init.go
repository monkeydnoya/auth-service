package mongodriver

import (
	"fmt"

	"github.com/monkeydnoya/hiraishin-auth/internal/data/repository"
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
	fmt.Println("[INFO] Auth repository for mongo driver init")
	auth := AuthDAO{
		DB: c.Database,
	}

	return auth, nil
}
