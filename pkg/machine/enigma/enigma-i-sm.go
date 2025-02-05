package enigma

import (
	"strings"

	"github.com/mrumyantsev/enigma-app/pkg/machine"
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/base"
	"github.com/mrumyantsev/enigma-app/pkg/machine/enigma/parts"
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

		switch strings.ToUpper(settings.Name) {
		case "II":
			rots[i] = parts.NewISMRotorII(settings.Position, settings.Ring)
		case "III":
			rots[i] = parts.NewISMRotorIII(settings.Position, settings.Ring)
		default:
			rots[i] = parts.NewISMRotorI(settings.Position, settings.Ring)
		}
	}

	sta := parts.NewIStator()

	pb := base.NewPlugboard(strings.ToUpper(pboard), base.CharsCount26)

	fil := base.NewFilter()

	tr := base.NewTranslator()

	return &EnigmaISM{
		base.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}

// EnigmaISMSpec returns the Enigma I Sondermaschine A-17401 S specification.
func EnigmaISMSpec() machine.MachineSpec {
	return machine.MachineSpec{
		Name: `Enigma I "Sondermaschine"`,
		Rotors: []machine.RotorSpec{
			{Name: "I", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "II", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "III", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
		},
		RotorsCount:     base.RotorsCount3,
		PlugboardsCount: base.PlugboardsCount1,
	}
}
