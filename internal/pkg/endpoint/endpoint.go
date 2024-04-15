package endpoint

import (
	"github.com/labstack/echo/v4"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/config"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/service"
)

type MachinesSpec interface {
	MachinesSpec(ctx echo.Context) error
}

type Plaintext interface {
	Plaintext(ctx echo.Context) error
}

type Endpoint struct {
	MachinesSpec MachinesSpec
	Plaintext    Plaintext
}

func New(cfg *config.Config, svc *service.Service) *Endpoint {
	return &Endpoint{
		MachinesSpec: NewMachinesSpecEndpoint(cfg, svc.MachinesSpec),
		Plaintext:    NewPlaintextEndpoint(cfg, svc.Ciphertext),
	}
}

func (e *Endpoint) InitRoutes(echo *echo.Echo) {
	echo.GET("/machines-spec", e.MachinesSpec.MachinesSpec)

	echo.POST("/plaintext", e.Plaintext.Plaintext)
}
