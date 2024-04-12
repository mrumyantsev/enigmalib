package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/config"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/handler"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/server"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/service"
	"github.com/mrumyantsev/logx/log"
)

type App struct {
	config  *config.Config
	handler *handler.Handler
	service *service.Service
	server  *server.Server
}

func New() *App {
	cfg := config.New()

	err := cfg.Init()
	if err != nil {
		log.Fatal("could not initialize application configuration", err)
	}

	mw := []echo.MiddlewareFunc{
		middleware.CORS(),
	}

	svc := service.New(cfg)

	hdl := handler.New(cfg, svc)

	srv := server.New(cfg, mw, hdl)

	return &App{
		config:  cfg,
		handler: hdl,
		service: svc,
		server:  srv,
	}
}

func (a *App) Run() {
	log.Info("service started")

	if err := a.server.Start(); err != nil {
		log.Fatal("could not start http server", err)
	}
}
