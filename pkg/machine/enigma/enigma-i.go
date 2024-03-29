package enigma

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
	enigmai "github.com/mrumyantsev/encryption-app/pkg/part/enigma-i"
)

type EnigmaI struct {
	*machine.Machine
}

func NewEnigmaI(reflector string, rotorSet [base.RotorsCount3]base.RotorSettings, pboard string) *EnigmaI {
	var refl base.Reflector

	switch reflector {
	case "B":
		refl = enigmai.NewReflectorB()
	case "C":
		refl = enigmai.NewReflectorC()
	default:
		refl = enigmai.NewReflectorA()
	}

	rots := make([]base.Rotor, base.RotorsCount3)
	var settings base.RotorSettings

	for i := 0; i < base.RotorsCount3; i++ {
		settings = rotorSet[i]

		switch settings.Name {
		case "II":
			rots[i] = enigmai.NewRotorII(settings.Pos, settings.RingPos)
		case "III":
			rots[i] = enigmai.NewRotorIII(settings.Pos, settings.RingPos)
		case "IV":
			rots[i] = enigmai.NewRotorIV(settings.Pos, settings.RingPos)
		case "V":
			rots[i] = enigmai.NewRotorV(settings.Pos, settings.RingPos)
		default:
			rots[i] = enigmai.NewRotorI(settings.Pos, settings.RingPos)
		}
	}

	sta := enigmai.NewStator()

	pb := machine.NewPlugboard(pboard, base.CharactersCount26)

	fil := machine.NewFilter()

	tr := machine.NewTranslator()

	return &EnigmaI{
		machine.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}
