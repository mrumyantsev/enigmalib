package endpoint

import (
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/config"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/service"
)

type Endpoint struct {
	config  *config.Config
	service *service.Service
}

func New(cfg *config.Config, svc *service.Service) *Endpoint {
	return &Endpoint{
		config:  cfg,
		service: svc,
	}
}
