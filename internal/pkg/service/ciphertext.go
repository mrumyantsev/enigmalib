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

type CiphertextService struct {
	config *config.Config
}

func NewCiphertextService(cfg *config.Config) *CiphertextService {
	return &CiphertextService{
		config: cfg,
	}
}

func (s *CiphertextService) Ciphertext(plaintext models.PlaintextMsg) (models.CiphertextMsg, error) {
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

	var ciphertext models.CiphertextMsg

	text, err := m.EncryptString(plaintext.Plaintext)
	if err != nil {
		return ciphertext, errlib.Wrap("could not encrypt via machine", err)
	}

	ciphertext.Ciphertext = text

	return ciphertext, nil
}
