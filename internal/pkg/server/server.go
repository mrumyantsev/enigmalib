package server

import (
	"github.com/labstack/echo/v4"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/config"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/handler"
	"github.com/mrumyantsev/cipher-machines-app/pkg/lib/e"
)

type Server struct {
	config *config.Config
	echo   *echo.Echo
}

func New(cfg *config.Config, mw []echo.MiddlewareFunc, hdl ...*handler.Handler) *Server {
	echo := echo.New()

	echo.HideBanner = true

	echo.Use(mw...)

	for _, handler := range hdl {
		handler.RegisterBy(echo)
	}

	return &Server{
		config: cfg,
		echo:   echo,
	}
}

func (s *Server) Start() error {
	if err := s.echo.Start(s.config.HttpServerListenAddress); err != nil {
		return e.Wrap("could not start echo server", err)
	}

	return nil
}
