package config

import (
	"io"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger = GenerateLog()

func GenerateLog() *zap.SugaredLogger {
	configFile, err := os.Open("config/logger-config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	defer configFile.Close()

	config, err := io.ReadAll(configFile)

	if err != nil {
		log.Fatal(err)
	}

	var cfg zap.Config
	if err := yaml.Unmarshal(config, &cfg); err != nil {
		log.Fatal(err)
	}

	os.MkdirAll("logs", 0764)
	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()
	return sugar
}
