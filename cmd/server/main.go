package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma"
	"github.com/mrumyantsev/logx/log"
)

const (
	listenAddr = "0.0.0.0:8080"
)

func main() {
	log.Info("service started")

	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.CORS())
	e.Use(errorLoggerMW)

	e.GET("/machines-spec.json", machinesSpec)
	e.POST("/encrypt-text", encryptText)

	go func() {
		err := e.Start(listenAddr)
		if err != nil {
			log.Fatal("could not start http server", err)
		}
	}()

	log.Info("http server started at " + listenAddr)

	time.Sleep(4 * time.Hour)
}

func errorLoggerMW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		err := next(ctx)
		if err != nil {
			log.Error("could not handle request", err)

			return err
		}

		return nil
	}
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

func encryptText(c echo.Context) error {
	req := c.Request()

	b, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("could not read all bytes from body: %v", err)
	}

	p := new(PlaintextMsg)

	err = json.Unmarshal(b, p)
	if err != nil {
		return fmt.Errorf("could not unmarshal json to PlaintextMsg struct: %v", err)
	}

	fmt.Println(p)

	return c.JSON(200, CiphertextMsg{Ciphertext: "XGWBSW"})
}

type PlaintextMsg struct {
	Plaintext string   `json:"plaintext"`
	Machine   string   `json:"machine"`
	Settings  Settings `json:"settings"`
}

type Settings struct {
	Reflector     string                  `json:"reflector"`
	RotorSettings []machine.RotorSettings `json:"rotorSettings"`
	Plugboard     string                  `json:"plugboard"`
}

type CiphertextMsg struct {
	Ciphertext string `json:"ciphertext"`
}
