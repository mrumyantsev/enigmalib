package endpoint

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/config"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/service"
)

type MachinesSpecEndpoint struct {
	config  *config.Config
	service service.MachinesSpec
}

func NewMachinesSpecEndpoint(cfg *config.Config, svc service.MachinesSpec) *MachinesSpecEndpoint {
	return &MachinesSpecEndpoint{
		config:  cfg,
		service: svc,
	}
}

func (e *MachinesSpecEndpoint) MachinesSpec(ctx echo.Context) error {
	spec := e.service.MachinesSpec()

	return ctx.JSON(http.StatusOK, spec)
}
