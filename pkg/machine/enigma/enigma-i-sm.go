package enigma

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
	enigmai "github.com/mrumyantsev/encryption-app/pkg/part/enigma-i"
	enigmaism "github.com/mrumyantsev/encryption-app/pkg/part/enigma-i-sm"
)

type EnigmaISondermaschine struct {
	*machine.Machine
}

func NewEnigmaISondermaschine(reflector string, rotorSet [base.RotorsCount3]base.RotorSettings, pboard string) *EnigmaISondermaschine {
	var refl base.Reflector

	switch reflector {
	default:
		refl = enigmaism.NewReflector()
	}

	rots := make([]base.Rotor, base.RotorsCount3)
	var settings base.RotorSettings

	for i := 0; i < base.RotorsCount3; i++ {
		settings = rotorSet[i]

		switch settings.Name {
		case "II":
			rots[i] = enigmaism.NewRotorII(settings.Pos, settings.RingPos)
		case "III":
			rots[i] = enigmaism.NewRotorIII(settings.Pos, settings.RingPos)
		default:
			rots[i] = enigmaism.NewRotorI(settings.Pos, settings.RingPos)
		}
	}

	sta := enigmai.NewStator()

	pb := machine.NewPlugboard(pboard, base.CharactersCount26)

	fil := machine.NewFilter()

	tr := machine.NewTranslator()

	return &EnigmaISondermaschine{
		machine.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}
