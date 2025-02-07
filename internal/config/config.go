package config

import (
	"context"
	"errors"

	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	DbFilePath string `env:"TODO_DBFILE"`
	Port       string `env:"TODO_PORT, default=7540"`
}

var errNoEnvDbFile error = errors.New("no specified environment variable TODO_DBFILE")

func LoadAppConfig() (AppConfig, error) {
	ctx := context.Background()

	var appConfig AppConfig
	if err := envconfig.Process(ctx, &appConfig); err != nil {
		return AppConfig{}, err
	}

	if appConfig.DbFilePath == "" {
		return AppConfig{}, errNoEnvDbFile
	}
	return appConfig, nil
}
