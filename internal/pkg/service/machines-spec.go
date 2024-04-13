package service

import (
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/config"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma"
)

type MachinesSpecService struct {
	config *config.Config
}

func NewMachinesSpecService(cfg *config.Config) *MachinesSpecService {
	return &MachinesSpecService{
		config: cfg,
	}
}

func (s *MachinesSpecService) MachinesSpec() machine.MachinesSpec {
	return machine.MachinesSpec{
		Machines: []machine.MachineSpec{
			enigma.EnigmaISpec(),
			enigma.EnigmaINESpec(),
			enigma.EnigmaISMSpec(),
			enigma.EnigmaM3Spec(),
			enigma.EnigmaM4SSpec(),
		},
	}
}
