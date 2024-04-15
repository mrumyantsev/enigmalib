package endpoint

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/config"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/models"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/service"
	"github.com/mrumyantsev/cipher-machines-app/pkg/lib/errlib"
)

type PlaintextEndpoint struct {
	config  *config.Config
	service service.Ciphertext
}

func NewPlaintextEndpoint(cfg *config.Config, svc service.Ciphertext) *PlaintextEndpoint {
	return &PlaintextEndpoint{
		config:  cfg,
		service: svc,
	}
}

func (e *PlaintextEndpoint) Plaintext(ctx echo.Context) error {
	req := ctx.Request()

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return errlib.Wrap("could not read all bytes from body", err)
	}
	defer func() { _ = req.Body.Close() }()

	var plaintext models.PlaintextMsg

	if err = json.Unmarshal(data, &plaintext); err != nil {
		return errlib.Wrap("could not unmarshal json to PlaintextMsg struct", err)
	}

	ciphertext, err := e.service.Ciphertext(plaintext)
	if err != nil {
		return errlib.Wrap("could not encrypt plaintext", err)
	}

	return ctx.JSON(http.StatusOK, ciphertext)
}
