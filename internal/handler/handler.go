package handler

import (
	"github.com/YulyaY/go_final_project.git/internal/config"
	"github.com/YulyaY/go_final_project.git/internal/domain/service"
)

type Handler struct {
	service *service.Service
	appCfg  config.AppConfig
}

var appConfig config.AppConfig

func New(service *service.Service, appCfg config.AppConfig) *Handler {
	appConfig = appCfg
	return &Handler{
		service: service,
		appCfg:  appCfg,
	}
}
