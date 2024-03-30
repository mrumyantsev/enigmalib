package enigma

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine"
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/parts"
)

type EnigmaINE struct {
	*base.Machine
}

func NewEnigmaINE(reflector string, rotorSet [base.RotorsCount3]machine.RotorSettings, pboard string) *EnigmaINE {
	refl := parts.NewINEReflector()

	rots := make([]base.Rotorer, base.RotorsCount3)
	var settings machine.RotorSettings

	for i := 0; i < base.RotorsCount3; i++ {
		settings = rotorSet[i]

		switch settings.Name {
		case "II":
			rots[i] = parts.NewINERotorII(settings.Pos, settings.RingPos)
		case "III":
			rots[i] = parts.NewINERotorIII(settings.Pos, settings.RingPos)
		case "IV":
			rots[i] = parts.NewINERotorIV(settings.Pos, settings.RingPos)
		case "V":
			rots[i] = parts.NewINERotorV(settings.Pos, settings.RingPos)
		default:
			rots[i] = parts.NewINERotorI(settings.Pos, settings.RingPos)
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

func EnigmaINESpec() machine.MachineSpec {
	return machine.MachineSpec{
		Name: "Enigma I \"Norenigma\"",
		Rotors: []machine.RotorSpec{
			{Name: "I", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "II", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "III", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "IV", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "V", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
		},
		RotorsCount:    base.RotorsCount3,
		IsHasPlugboard: true,
	}
}
