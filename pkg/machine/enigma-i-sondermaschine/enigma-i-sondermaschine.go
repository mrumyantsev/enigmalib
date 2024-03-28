package enigmaisondermaschine

import (
	"github.com/mrumyantsev/enigma/pkg/base"
	"github.com/mrumyantsev/enigma/pkg/machine"
)

const (
	CharactersCount = 26
	RotorsCount     = 3
)

type EnigmaISondermaschine struct {
	*machine.Machine
}

func NewEnigmaISondermaschine(reflector string, rotorSet [RotorsCount]base.RotorSettings, stator string, pboard string) *EnigmaISondermaschine {
	var refl base.Reflector

	switch reflector {
	default:
		refl = NewReflector()
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

	return &EnigmaISondermaschine{
		machine.NewMachine(refl, rots, sta, pb, fil, tr),
	}
}
