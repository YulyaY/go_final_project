package config

import (
	"context"
	"errors"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	DbFilePath  string `env:"TODO_DBFILE, default=scheduler.db"`
	Port        string `env:"TODO_PORT, default=7540"`
	AppPassword string `env:"TODO_PASSWORD"`
	Secret      string `env:"TODO_SECRET"`
	DbName      string `env:"TODO_DBNAMEPGS, default=scheduler"`
	UserNamePG  string `env:"TODO_USERNAMEPG, default=postgres"`
	HostPG      string `env:"TODO_HOSTPG, default=127.0.0.1"`
	PortPG      string `env:"TODO_PORTPG, default=5432"`
	DbNamePG    string `env:"TODO_DBNAMEPG, default=scheduler"`
	PasswordPG  string `env:"TODO_PASSWORDPG"`
}

var errNoEnvDbFile error = errors.New("no specified environment variable TODO_DBFILE")
var errNoEnvSecret error = errors.New("no specified environment variable TODO_SECRET")

func LoadAppConfig() (AppConfig, error) {
	ctx := context.Background()
	err := godotenv.Load()
	if err == nil {
		log.Println("env variables loaded from .env file")
	}

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
