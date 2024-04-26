package server

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/mrumyantsev/enigma-app/internal/pkg/config"
	"github.com/mrumyantsev/enigma-app/internal/pkg/endpoint"
	"github.com/mrumyantsev/enigma-app/internal/pkg/server"
	"github.com/mrumyantsev/enigma-app/internal/pkg/service"
	"github.com/mrumyantsev/enigma-app/pkg/lib/errlib"
	"github.com/mrumyantsev/logx/log"
)

type App struct {
	config   *config.Config
	service  *service.Service
	endpoint *endpoint.Endpoint
	server   *server.Server
}

func New() (*App, error) {
	cfg := config.New()

	if err := cfg.Init(); err != nil {
		return nil, errlib.Wrap("could not initialize configuration", err)
	}

	service := service.New(cfg)

	endpoint := endpoint.New(cfg, service)

	mwCors := middleware.CORS()

	server := server.New(cfg, endpoint, mwCors)

	return &App{
		config:   cfg,
		service:  service,
		endpoint: endpoint,
		server:   server,
	}, nil
}

func (a *App) Run() error {
	log.Info("service started")

	if err := a.server.Start(); err != nil {
		return errlib.Wrap("could not start http server", err)
	}

	return nil
}
