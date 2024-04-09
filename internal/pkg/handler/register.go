package handler

import "github.com/labstack/echo/v4"

func (h *Handler) RegisterBy(echo *echo.Echo) {
	echo.GET("/machines-spec.json", h.MachinesSpec)

	echo.POST("/encrypt-text", h.EncryptText)
}
