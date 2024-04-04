package enigma

import (
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/base"
	"github.com/mrumyantsev/cipher-machines-app/pkg/machine/enigma/parts"
)

// An Enigma M4 "Shark" model.
type EnigmaM4S struct {
	*base.Machine
}

// NewEnigmaM4S creates the Enigma M4 "Shark" model.
func NewEnigmaM4S(reflector string, rotorSet [base.RotorsCount4]machine.RotorSettings, pboard string) *EnigmaM4S {
	var refl base.Reflectorer

	switch reflector {
	case "C THIN":
		refl = parts.NewM4SReflectorCThin()
	default:
		refl = parts.NewM4SReflectorBThin()
	}

	rots := make([]base.Rotorer, base.RotorsCount4)
	var settings machine.RotorSettings

	for i := 0; i < base.RotorsCount4; i++ {
		settings = rotorSet[i]

		switch settings.Name {
		case "GAMMA":
			rots[i] = parts.NewM4SRotorGamma(settings.Position, settings.Ring)
		case "I":
			rots[i] = parts.NewIRotorI(settings.Position, settings.Ring)
		case "II":
			rots[i] = parts.NewIRotorII(settings.Position, settings.Ring)
		case "III":
			rots[i] = parts.NewIRotorIII(settings.Position, settings.Ring)
		case "IV":
			rots[i] = parts.NewIRotorIV(settings.Position, settings.Ring)
		case "V":
			rots[i] = parts.NewIRotorV(settings.Position, settings.Ring)
		case "VI":
			rots[i] = parts.NewM3RotorVI(settings.Position, settings.Ring)
		case "VII":
			rots[i] = parts.NewM3RotorVII(settings.Position, settings.Ring)
		case "VIII":
			rots[i] = parts.NewM3RotorVIII(settings.Position, settings.Ring)
		default:
			rots[i] = parts.NewM4SRotorBeta(settings.Position, settings.Ring)
		}
	}

	sta := parts.NewIStator()

	pb := base.NewPlugboard(pboard, base.CharsCount26)

	fil := base.NewFilter()

	tr := base.NewTranslator()

	return &EnigmaM4S{
		base.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}

// EnigmaM4SSpec returns the Enigma M4 "Shark" specification.
func EnigmaM4SSpec() machine.MachineSpec {
	return machine.MachineSpec{
		Name: "Enigma M4 \"Shark\"",
		Reflectors: []machine.ReflectorSpec{
			{Name: "B Thin"},
			{Name: "C Thin"},
		},
		Rotors: []machine.RotorSpec{
			{Name: "Beta", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
			{Name: "Gamma", Positions: base.CharsCount26, RingPositions: base.CharsCount26},
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
		RotorsCount:     base.RotorsCount4,
		PlugboardsCount: base.PlugboardsCount1,
	}
}
