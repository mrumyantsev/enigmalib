package enigma

import (
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/parts"
)

// An Enigma I Sondermaschine A-17401 S model.
type EnigmaISM struct {
	*base.Machine
}

// NewEnigmaISM creates the Enigma I Sondermaschine A-17401 S model.
func NewEnigmaISM(reflector string, rotorSet [base.RotorsCount3]machine.RotorSettings, pboard string) *EnigmaISM {
	refl := parts.NewISMReflector()

	rots := make([]base.Rotorer, base.RotorsCount3)
	var settings machine.RotorSettings

	for i := 0; i < base.RotorsCount3; i++ {
		settings = rotorSet[i]

		switch settings.Name {
		case "II":
			rots[i] = parts.NewISMRotorII(settings.Pos, settings.RingPos)
		case "III":
			rots[i] = parts.NewISMRotorIII(settings.Pos, settings.RingPos)
		default:
			rots[i] = parts.NewISMRotorI(settings.Pos, settings.RingPos)
		}
	}

	sta := parts.NewIStator()

	pb := base.NewPlugboard(pboard, base.CharsCount26)

	fil := base.NewFilter()

	tr := base.NewTranslator()

	return &EnigmaISM{
		base.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}

// EnigmaISMSpec returns the Enigma I Sondermaschine A-17401 S specification.
func EnigmaISMSpec() machine.MachineSpec {
	return machine.MachineSpec{
		Name: "Enigma I \"Sondermaschine\"",
		Rotors: []machine.RotorSpec{
			{Name: "I", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "II", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "III", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
		},
		ReflectorsCount: base.ReflectorsCount1,
		RotorsCount:     base.RotorsCount3,
		PlugboardsCount: base.PlugboardsCount1,
	}
}
