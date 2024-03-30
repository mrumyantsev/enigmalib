package enigma

import (
	"github.com/mrumyantsev/encryption-app/pkg/machine"
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine/enigma/parts"
)

type EnigmaISM struct {
	*base.Machine
}

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

func EnigmaISMSpec() machine.MachineSpec {
	return machine.MachineSpec{
		Name: "Enigma I \"Sondermaschine\"",
		Rotors: []machine.RotorSpec{
			{Name: "I", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "II", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
			{Name: "III", Poses: base.CharsCount26, RingPoses: base.CharsCount26},
		},
		RotorsCount:    base.RotorsCount3,
		IsHasPlugboard: true,
	}
}
