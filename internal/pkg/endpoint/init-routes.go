package endpoint

import "github.com/labstack/echo/v4"

func (e *Endpoint) InitRoutes(echo *echo.Echo) {
	echo.GET("/machines-spec", e.MachinesSpec)

	echo.POST("/plaintext", e.Plaintext)
}
