package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma"
	"github.com/mrumyantsev/logx/log"
)

func main() {
	log.Info("service started")

	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.CORS())

	e.GET("/machines-spec.json", machinesSpec)

	e.Logger.Fatal(e.Start(":8080"))
}

func machinesSpec(c echo.Context) error {
	spec := machine.MachinesSpec{
		Machines: []machine.MachineSpec{
			enigma.EnigmaISpec(),
			enigma.EnigmaINESpec(),
			enigma.EnigmaISMSpec(),
			enigma.EnigmaM3Spec(),
			enigma.EnigmaM4SSpec(),
		},
	}

	return c.JSON(200, spec)
}