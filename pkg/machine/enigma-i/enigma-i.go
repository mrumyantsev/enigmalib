package enigmai

import (
	"github.com/mrumyantsev/encryption-app/pkg/base"
	"github.com/mrumyantsev/encryption-app/pkg/machine"
)

const (
	CharactersCount = 26
	RotorsCount     = 3
)

type EnigmaI struct {
	*machine.Machine
}

func NewEnigmaI(reflector string, rotorSet [RotorsCount]base.RotorSettings, stator string, pboard string) *EnigmaI {
	var refl base.Reflector

	switch reflector {
	case "B":
		refl = NewReflectorB()
	case "C":
		refl = NewReflectorC()
	default:
		refl = NewReflectorA()
	}

	rots := make([]base.Rotor, RotorsCount)

	for i := 0; i < RotorsCount; i++ {
		switch rotorSet[i].Name {
		case "II":
			rots[i] = NewRotorII(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "III":
			rots[i] = NewRotorIII(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "IV":
			rots[i] = NewRotorIV(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		case "V":
			rots[i] = NewRotorV(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		default:
			rots[i] = NewRotorI(
				rotorSet[i].Pos,
				rotorSet[i].RingPos,
			)
		}
	}

	var sta base.Stator

	switch stator {
	default:
		sta = NewStator()
	}

	pb := machine.NewPlugboard(pboard, CharactersCount)

	fil := machine.NewFilter()

	tr := machine.NewTranslator()

	return &EnigmaI{
		machine.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}
