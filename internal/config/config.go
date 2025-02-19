package config

import (
	"context"
	"errors"

	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	DbFilePath  string `env:"TODO_DBFILE, default=scheduler.db"`
	Port        string `env:"TODO_PORT, default=7540"`
	AppPassword string `env:"TODO_PASSWORD"`
	Secret      string `env:"TODO_SECRET"`
}

var errNoEnvDbFile error = errors.New("no specified environment variable TODO_DBFILE")
var errNoEnvSecret error = errors.New("no specified environment variable TODO_SECRET")

func LoadAppConfig() (AppConfig, error) {
	ctx := context.Background()

	var appConfig AppConfig
	if err := envconfig.Process(ctx, &appConfig); err != nil {
		return AppConfig{}, err
	}

	if appConfig.DbFilePath == "" {
		return AppConfig{}, errNoEnvDbFile
	}

	if appConfig.AppPassword != "" && appConfig.Secret == "" {
		return AppConfig{}, errNoEnvSecret
	}
	return appConfig, nil
}

func (cfg *AppConfig) IsPasswordSet() bool {
	return len(cfg.AppPassword) > 0
}
