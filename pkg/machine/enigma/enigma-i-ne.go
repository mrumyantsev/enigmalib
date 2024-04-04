package enigma

import (
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/parts"
)

// An Enigma I Norway Enigma model.
type EnigmaINE struct {
	*base.Machine
}

// NewEnigmaINE creates the Enigma I Norway Enigma model.
func NewEnigmaINE(reflector string, rotorSet [base.RotorsCount3]machine.RotorSettings, pboard string) *EnigmaINE {
	refl := parts.NewINEReflector()

	rots := make([]base.Rotorer, base.RotorsCount3)
	var settings machine.RotorSettings

	for i := 0; i < base.RotorsCount3; i++ {
		settings = rotorSet[i]

		switch settings.Name {
		case "II":
			rots[i] = parts.NewINERotorII(settings.Position, settings.Ring)
		case "III":
			rots[i] = parts.NewINERotorIII(settings.Position, settings.Ring)
		case "IV":
			rots[i] = parts.NewINERotorIV(settings.Position, settings.Ring)
		case "V":
			rots[i] = parts.NewINERotorV(settings.Position, settings.Ring)
		default:
			rots[i] = parts.NewINERotorI(settings.Position, settings.Ring)
		}
	}

	sta := parts.NewIStator()

	pb := base.NewPlugboard(pboard, base.CharsCount26)

	fil := base.NewFilter()

	tr := base.NewTranslator()

	return &EnigmaINE{
		base.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}

// EnigmaINESpec returns the Enigma I Norway Enigma specification.
func EnigmaINESpec() machine.MachineSpec {
	return machine.MachineSpec{
		Name: "Enigma I \"Norenigma\"",
		Rotors: []machine.RotorSpec{
			{Name: "I", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "II", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "III", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "IV", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "V", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
		},
		RotorsCount:     base.RotorsCount3,
		PlugboardsCount: base.PlugboardsCount1,
	}
}
