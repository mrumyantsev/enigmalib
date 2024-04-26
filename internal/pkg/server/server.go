package server

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/mrumyantsev/enigma-app/internal/pkg/config"
	"github.com/mrumyantsev/enigma-app/internal/pkg/endpoint"
	"github.com/mrumyantsev/enigma-app/pkg/lib/errlib"
)

type Server struct {
	config *config.Config
	echo   *echo.Echo
}

func New(cfg *config.Config, ep *endpoint.Endpoint, mw ...echo.MiddlewareFunc) *Server {
	echo := echo.New()

	echo.HideBanner = true

	ep.InitRoutes(echo)

	echo.Use(mw...)

	return &Server{
		config: cfg,
		echo:   echo,
	}
}

func (s *Server) Start() error {
	listenAddr := s.config.HttpServerListenIp + ":" + s.config.HttpServerListenPort

	if err := s.echo.Start(listenAddr); err != nil {
		return errlib.Wrap("could not start http server", err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.echo.Shutdown(ctx); err != nil {
		return errlib.Wrap("could not shutdown http server", err)
	}

	return nil
}
