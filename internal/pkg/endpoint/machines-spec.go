package endpoint

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (e *Endpoint) MachinesSpec(ctx echo.Context) error {
	spec := e.service.MachinesSpec.MachinesSpec()

	return ctx.JSON(http.StatusOK, spec)
}
