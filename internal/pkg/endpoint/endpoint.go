package endpoint

import (
	"github.com/labstack/echo/v4"
	"github.com/mrumyantsev/enigma-app/internal/pkg/config"
	"github.com/mrumyantsev/enigma-app/internal/pkg/service"
)

type MachinesSpec interface {
	MachinesSpec(ctx echo.Context) error
}

type Encryption interface {
	Encryption(ctx echo.Context) error
}

type Endpoint struct {
	MachinesSpec MachinesSpec
	Encryption   Encryption
}

func New(cfg *config.Config, svc *service.Service) *Endpoint {
	return &Endpoint{
		MachinesSpec: NewMachinesSpecEndpoint(cfg, svc.MachinesSpec),
		Encryption:   NewEncryptionEndpoint(cfg, svc.Encryption),
	}
}

func (e *Endpoint) InitRoutes(echo *echo.Echo) {
	echo.GET("/machines-spec", e.MachinesSpec.MachinesSpec)

	echo.POST("/encryption", e.Encryption.Encryption)
}
