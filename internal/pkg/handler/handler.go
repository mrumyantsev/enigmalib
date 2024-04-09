package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/config"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/models"
	"github.com/mrumyantsev/cipher-machines-app/pkg/lib/e"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma"
)

type Service interface {
	Ciphertext(plaintext models.PlaintextMsg) (models.CiphertextMsg, error)
}

type Handler struct {
	config  *config.Config
	service Service
}

func New(cfg *config.Config, s Service) *Handler {
	return &Handler{
		config:  cfg,
		service: s,
	}
}

func (h *Handler) MachinesSpec(ctx echo.Context) error {
	spec := machine.MachinesSpec{
		Machines: []machine.MachineSpec{
			enigma.EnigmaISpec(),
			enigma.EnigmaINESpec(),
			enigma.EnigmaISMSpec(),
			enigma.EnigmaM3Spec(),
			enigma.EnigmaM4SSpec(),
		},
	}

	return ctx.JSON(http.StatusOK, spec)
}

func (h *Handler) EncryptText(ctx echo.Context) error {
	req := ctx.Request()

	b, err := io.ReadAll(req.Body)
	if err != nil {
		return e.Wrap("could not read all bytes from body", err)
	}

	var p models.PlaintextMsg

	if err = json.Unmarshal(b, &p); err != nil {
		return e.Wrap("could not unmarshal json to PlaintextMsg struct", err)
	}

	ciphertext, err := h.service.Ciphertext(p)
	if err != nil {
		return e.Wrap("could not encrypt plaintext", err)
	}

	return ctx.JSON(http.StatusOK, ciphertext)
}
