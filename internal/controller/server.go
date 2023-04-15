package controller

import (
	"github.com/monkeydnoya/hiraishin-auth/internal/controller/fiber/handler"
	configuration "github.com/monkeydnoya/hiraishin-auth/pkg/config"
)

func StartServer(server handler.Server) {
	configuration.Logger.Infow("Starting Auth Service")

	port := ":" + configuration.Config("SERVICE_PORT")
	server.App.Listen(port)
}
