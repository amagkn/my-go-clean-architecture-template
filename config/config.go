package config

import (
	"github.com/amagkn/my-go-clean-architecture-template/pkg/base_errors"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/http_server"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/logger"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/postgres"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

/*
Config holds environment variable data.
Initialization steps:

 1. Use godotenv to read .env file
 2. Use envconfig to map variables to the struct
*/
type Config struct {
	App      App
	Logger   logger.Config
	HTTP     http_server.Config
	Postgres postgres.Config
}

func New() (Config, error) {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		return config, base_errors.WithPath("godotenv.Load", err)
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return config, base_errors.WithPath("envconfig.Process", err)
	}

	return config, nil
}
