package enigma

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
	enigmai "github.com/mrumyantsev/encryption-app/pkg/part/enigma-i"
	enigmam3 "github.com/mrumyantsev/encryption-app/pkg/part/enigma-m3"
)

type EnigmaM3 struct {
	*machine.Machine
}

func NewEnigmaM3(reflector string, rotorSet [base.RotorsCount3]base.RotorSettings, pboard string) *EnigmaM3 {
	var refl base.Reflector

	switch reflector {
	case "C":
		refl = enigmai.NewReflectorC()
	default:
		refl = enigmai.NewReflectorB()
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
		case "VI":
			rots[i] = enigmam3.NewRotorVI(settings.Pos, settings.RingPos)
		case "VII":
			rots[i] = enigmam3.NewRotorVII(settings.Pos, settings.RingPos)
		case "VIII":
			rots[i] = enigmam3.NewRotorVIII(settings.Pos, settings.RingPos)
		default:
			rots[i] = enigmai.NewRotorI(settings.Pos, settings.RingPos)
		}
	}

	sta := enigmai.NewStator()

	pb := machine.NewPlugboard(pboard, base.CharactersCount26)

	fil := machine.NewFilter()

	tr := machine.NewTranslator()

	return &EnigmaM3{
		machine.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}
