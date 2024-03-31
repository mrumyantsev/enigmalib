package enigma

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine"
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/parts"
)

// An Enigma M3 model.
type EnigmaM3 struct {
	*base.Machine
}

// NewEnigmaM3 creates the Enigma M3 model.
func NewEnigmaM3(reflector string, rotorSet [base.RotorsCount3]machine.RotorSettings, pboard string) *EnigmaM3 {
	var refl base.Reflectorer

	switch reflector {
	case "C":
		refl = parts.NewIReflectorC()
	default:
		refl = parts.NewIReflectorB()
	}

	rots := make([]base.Rotorer, base.RotorsCount3)
	var settings machine.RotorSettings

	for i := 0; i < base.RotorsCount3; i++ {
		settings = rotorSet[i]

		switch settings.Name {
		case "II":
			rots[i] = parts.NewIRotorII(settings.Pos, settings.RingPos)
		case "III":
			rots[i] = parts.NewIRotorIII(settings.Pos, settings.RingPos)
		case "IV":
			rots[i] = parts.NewIRotorIV(settings.Pos, settings.RingPos)
		case "V":
			rots[i] = parts.NewIRotorV(settings.Pos, settings.RingPos)
		case "VI":
			rots[i] = parts.NewM3RotorVI(settings.Pos, settings.RingPos)
		case "VII":
			rots[i] = parts.NewM3RotorVII(settings.Pos, settings.RingPos)
		case "VIII":
			rots[i] = parts.NewM3RotorVIII(settings.Pos, settings.RingPos)
		default:
			rots[i] = parts.NewIRotorI(settings.Pos, settings.RingPos)
		}
	}

	sta := parts.NewIStator()

	pb := base.NewPlugboard(pboard, base.CharsCount26)

	fil := base.NewFilter()

	tr := base.NewTranslator()

	return &EnigmaM3{
		base.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}

// EnigmaM3Spec returns the Enigma M3 specification.
func EnigmaM3Spec() machine.MachineSpec {
	return machine.MachineSpec{
		Name: "Enigma M3",
		Reflectors: []machine.ReflectorSpec{
			{Name: "B"},
			{Name: "C"},
		},
		Rotors: []machine.RotorSpec{
			{Name: "I", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "II", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "III", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "IV", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "V", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "VI", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "VII", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "VIII", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
		},
		ReflectorsCount: base.ReflectorsCount1,
		RotorsCount:     base.RotorsCount3,
		PlugboardsCount: base.PlugboardsCount1,
	}
}
