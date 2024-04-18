package service

import (
	"strings"

	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/config"
	"github.com/mrumyantsev/cipher-machines-app/internal/pkg/models"
	"github.com/mrumyantsev/cipher-machines-app/pkg/lib/errlib"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
)

type EncryptionService struct {
	config *config.Config
}

func NewEncryptionService(cfg *config.Config) *EncryptionService {
	return &EncryptionService{
		config: cfg,
	}
}

func (s *EncryptionService) Encryption(plaintext models.PlaintextMsg) (models.CiphertextMsg, error) {
	var m machine.Machiner

	machineName := strings.ToLower(plaintext.Machine)

	switch machineName {
	case `enigma i "norenigma"`:
		m = enigma.NewEnigmaINE(
			plaintext.Settings.Reflector,
			[base.RotorsCount3]machine.RotorSettings(plaintext.Settings.RotorSettings),
			plaintext.Settings.Plugboard,
		)
	case `enigma i "sondermaschine"`:
		m = enigma.NewEnigmaISM(
			plaintext.Settings.Reflector,
			[base.RotorsCount3]machine.RotorSettings(plaintext.Settings.RotorSettings),
			plaintext.Settings.Plugboard,
		)
	case "enigma m3":
		m = enigma.NewEnigmaM3(
			plaintext.Settings.Reflector,
			[base.RotorsCount3]machine.RotorSettings(plaintext.Settings.RotorSettings),
			plaintext.Settings.Plugboard,
		)
	case `enigma m4 "shark"`:
		m = enigma.NewEnigmaM4S(
			plaintext.Settings.Reflector,
			[base.RotorsCount4]machine.RotorSettings(plaintext.Settings.RotorSettings),
			plaintext.Settings.Plugboard,
		)
	default:
		m = enigma.NewEnigmaI(
			plaintext.Settings.Reflector,
			[base.RotorsCount3]machine.RotorSettings(plaintext.Settings.RotorSettings),
			plaintext.Settings.Plugboard,
		)
	}

	var ciphertext models.CiphertextMsg

	text, err := m.EncryptString(plaintext.Plaintext)
	if err != nil {
		return ciphertext, errlib.Wrap("could not encrypt via machine", err)
	}

	ciphertext.Ciphertext = text

	return ciphertext, nil
}
