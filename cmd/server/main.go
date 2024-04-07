package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
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

	ct, err := ciphertext(*p)
	if err != nil {
		return fmt.Errorf("could not encrypt plaintext: %v", err)
	}

	return c.JSON(200, ct)
}

func ciphertext(plaintext PlaintextMsg) (CiphertextMsg, error) {
	var m machine.Machiner

	if strings.EqualFold(plaintext.Machine, `Enigma I "Norenigma"`) {
		m = enigma.NewEnigmaINE(
			plaintext.Settings.Reflector,
			[base.RotorsCount3]machine.RotorSettings(plaintext.Settings.RotorSettings),
			plaintext.Settings.Plugboard,
		)
	} else if strings.EqualFold(plaintext.Machine, `Enigma I "Sondermaschine"`) {
		m = enigma.NewEnigmaISM(
			plaintext.Settings.Reflector,
			[base.RotorsCount3]machine.RotorSettings(plaintext.Settings.RotorSettings),
			plaintext.Settings.Plugboard,
		)
	} else if strings.EqualFold(plaintext.Machine, "Enigma M3") {
		m = enigma.NewEnigmaM3(
			plaintext.Settings.Reflector,
			[base.RotorsCount3]machine.RotorSettings(plaintext.Settings.RotorSettings),
			plaintext.Settings.Plugboard,
		)
	} else if strings.EqualFold(plaintext.Machine, `Enigma M4 "Shark"`) {
		m = enigma.NewEnigmaM4S(
			plaintext.Settings.Reflector,
			[base.RotorsCount4]machine.RotorSettings(plaintext.Settings.RotorSettings),
			plaintext.Settings.Plugboard,
		)
	} else {
		m = enigma.NewEnigmaI(
			plaintext.Settings.Reflector,
			[base.RotorsCount3]machine.RotorSettings(plaintext.Settings.RotorSettings),
			plaintext.Settings.Plugboard,
		)
	}

	ciphertext := CiphertextMsg{}

	text, err := m.EncryptString(plaintext.Plaintext)
	if err != nil {
		return ciphertext, fmt.Errorf("could not encrypt via machine: %v", err)
	}

	ciphertext.Ciphertext = text

	return ciphertext, nil
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
