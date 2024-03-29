package enigma

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
	enigmai "github.com/mrumyantsev/encryption-app/pkg/part/enigma-i"
	enigmaine "github.com/mrumyantsev/encryption-app/pkg/part/enigma-i-ne"
)

type EnigmaINorenigma struct {
	*machine.Machine
}

func NewEnigmaINorenigma(reflector string, rotorSet [base.RotorsCount3]base.RotorSettings, pboard string) *EnigmaINorenigma {
	refl := enigmaine.NewReflector()

	rots := make([]base.Rotor, base.RotorsCount3)
	var settings base.RotorSettings

	for i := 0; i < base.RotorsCount3; i++ {
		settings = rotorSet[i]

		switch settings.Name {
		case "II":
			rots[i] = enigmaine.NewRotorII(settings.Pos, settings.RingPos)
		case "III":
			rots[i] = enigmaine.NewRotorIII(settings.Pos, settings.RingPos)
		case "IV":
			rots[i] = enigmaine.NewRotorIV(settings.Pos, settings.RingPos)
		case "V":
			rots[i] = enigmaine.NewRotorV(settings.Pos, settings.RingPos)
		default:
			rots[i] = enigmaine.NewRotorI(settings.Pos, settings.RingPos)
		}
	}

	sta := enigmai.NewStator()

	pb := machine.NewPlugboard(pboard, base.CharactersCount26)

	fil := machine.NewFilter()

	tr := machine.NewTranslator()

	return &EnigmaINorenigma{
		machine.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}

func EnigmaINorenigmaSpec() base.MachineSpec {
	return base.MachineSpec{
		Name: "Enigma I \"Norenigma\"",
		Rotors: []base.RotorSpec{
			{Name: "I", Poses: base.CharactersCount26, RingPoses: base.CharactersCount26},
			{Name: "II", Poses: base.CharactersCount26, RingPoses: base.CharactersCount26},
			{Name: "III", Poses: base.CharactersCount26, RingPoses: base.CharactersCount26},
			{Name: "IV", Poses: base.CharactersCount26, RingPoses: base.CharactersCount26},
			{Name: "V", Poses: base.CharactersCount26, RingPoses: base.CharactersCount26},
		},
		RotorsCount:    base.RotorsCount3,
		IsHasPlugboard: true,
	}
}
