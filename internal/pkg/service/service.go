package service

import (
	"github.com/mrumyantsev/enigma-app/internal/pkg/config"
	"github.com/mrumyantsev/enigma-app/internal/pkg/models"
	"github.com/mrumyantsev/enigma-app/pkg/machine"
)

type MachinesSpec interface {
	MachinesSpec() machine.MachinesSpec
}

type Encryption interface {
	Encryption(plaintext models.PlaintextMsg) (models.CiphertextMsg, error)
}

type Service struct {
	MachinesSpec MachinesSpec
	Encryption   Encryption
}

func New(cfg *config.Config) *Service {
	return &Service{
		MachinesSpec: NewMachinesSpecService(cfg),
		Encryption:   NewEncryptionService(cfg),
	}
}
