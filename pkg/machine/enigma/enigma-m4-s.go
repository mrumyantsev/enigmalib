package enigma

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine"
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/parts"
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
			rots[i] = parts.NewM4SRotorGamma(settings.Pos, settings.RingPos)
		case "I":
			rots[i] = parts.NewIRotorI(settings.Pos, settings.RingPos)
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
			rots[i] = parts.NewM4SRotorBeta(settings.Pos, settings.RingPos)
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
			{Name: "Beta", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "Gamma", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "I", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "II", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "III", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "IV", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "V", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "VI", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "VII", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "VIII", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
		},
		RotorsCount:    base.RotorsCount4,
		IsHasPlugboard: true,
	}
}
