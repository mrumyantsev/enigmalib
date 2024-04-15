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

type EncryptionEndpoint struct {
	config  *config.Config
	service service.Encryption
}

func NewEncryptionEndpoint(cfg *config.Config, svc service.Encryption) *EncryptionEndpoint {
	return &EncryptionEndpoint{
		config:  cfg,
		service: svc,
	}
}

func (e *EncryptionEndpoint) Encryption(ctx echo.Context) error {
	req := ctx.Request()

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return errlib.Wrap("could not read all bytes from body", err)
	}
	defer func() { _ = req.Body.Close() }()

	var plaintext models.PlaintextMsg

	if err = json.Unmarshal(data, &plaintext); err != nil {
		return errlib.Wrap("could not unmarshal json", err)
	}

	ciphertext, err := e.service.Encryption(plaintext)
	if err != nil {
		return errlib.Wrap("could not encrypt plaintext", err)
	}

	return ctx.JSON(http.StatusOK, ciphertext)
}
