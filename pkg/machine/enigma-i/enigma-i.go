package enigmai

import (
	"github.com/mrumyantsev/enigma/pkg/base"
	"github.com/mrumyantsev/enigma/pkg/machine"
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
	j := RotorsCount - 1
	var reversedItem base.RotorSettings

	for i := 0; i < RotorsCount; i++ {
		reversedItem = rotorSet[j]

		switch reversedItem.Name {
		case "II":
			rots[i] = NewRotorII(
				byte(rotorSet[i].Pos),
				byte(rotorSet[i].RingPos),
			)
		case "III":
			rots[i] = NewRotorIII(
				byte(rotorSet[i].Pos),
				byte(rotorSet[i].RingPos),
			)
		case "IV":
			rots[i] = NewRotorIV(
				byte(rotorSet[i].Pos),
				byte(rotorSet[i].RingPos),
			)
		case "V":
			rots[i] = NewRotorV(
				byte(rotorSet[i].Pos),
				byte(rotorSet[i].RingPos),
			)
		default:
			rots[i] = NewRotorI(
				byte(rotorSet[i].Pos),
				byte(rotorSet[i].RingPos),
			)
		}

		j--
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
