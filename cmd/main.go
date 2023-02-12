package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/monkeydnoya/hiraishin-auth/internal/data/config/mongodb"
	"github.com/monkeydnoya/hiraishin-auth/internal/data/repository/mongodriver"
)

func main() {
	godotenv.Load(".env")
	config := &mongodb.Config{
		Host:   os.Getenv("MONGO_URI"),
		Port:   os.Getenv("PORT"),
		DBName: os.Getenv("DB"),
		AppEnv: os.Getenv("APP_ENV"),
	}

	mongoconnect, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	}

	authDAO, err := mongodriver.Config{Client: mongoconnect.Client, Database: mongoconnect.DbConnection}.Init() // -> Interface 

	if err != nil {
		fmt.Println(err)
	}
	
	// Service init
	service := service.NewService{
		DAO: authDAO,
	}

}
